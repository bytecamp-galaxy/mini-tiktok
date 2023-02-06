package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/relation"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	q := query.Use(mysql.DB)
	err = q.Transaction(func(tx *query.Query) error {
		switch req.ActionType {
		case 1:
			{
				// 添加关注数据
				id := snowflake.Generate()
				err := tx.FollowRelation.WithContext(ctx).Create(&model.FollowRelation{
					ID:       id,
					UserID:   req.UserId,
					ToUserID: req.ToUserId,
				})
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				// 修改 FollowingCount 和 FollowerCount
				u := tx.User
				result, err := u.WithContext(ctx).Where(u.ID.Eq(req.UserId)).Update(u.FollowingCount, u.FollowingCount.Add(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
				}

				result, err = u.WithContext(ctx).Where(u.ID.Eq(req.ToUserId)).Update(u.FollowerCount, u.FollowerCount.Add(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
				}

				// update redis if exists
				exist, err := redis.FollowKeyExist(ctx, req.UserId)
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
				}
				if exist {
					count, err := redis.FollowKeyAdd(ctx, req.UserId, req.ToUserId)
					if err != nil {
						return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
					}
					if count != 1 {
						return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis sadd error")
					}
				} else {
					// do nothing
				}
			}
		case 2:
			{
				// 删除关注数据
				r := tx.FollowRelation
				result, err := r.WithContext(ctx).Where(r.UserID.Eq(req.UserId), r.ToUserID.Eq(req.ToUserId)).Delete()
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected == 0 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "nonexistent relation")
				}

				// 修改 FollowingCount 和 FollowerCount
				u := tx.User
				result, err = u.WithContext(ctx).Where(u.ID.Eq(req.UserId)).Update(u.FollowingCount, u.FollowingCount.Sub(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
				}

				result, err = u.WithContext(ctx).Where(u.ID.Eq(req.ToUserId)).Update(u.FollowerCount, u.FollowerCount.Sub(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
				}

				// update redis if exists
				exist, err := redis.FollowKeyExist(ctx, req.UserId)
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
				}
				if exist {
					count, err := redis.FollowKeyRem(ctx, req.UserId, req.ToUserId)
					if err != nil {
						return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
					}
					if count != 1 {
						return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis srem error")
					}
				} else {
					// do nothing
				}
			}
		default:
			return kerrors.NewBizStatusError(int32(errno.ErrUnknown), "unknown action type")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &relation.RelationActionResponse{}, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	r := query.FollowRelation
	u := query.User

	relList, err := r.WithContext(ctx).Preload(r.ToUser).Where(r.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	view, err := u.WithContext(ctx).Where(u.ID.Eq(req.UserViewId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i], err = convert.UserConverterORM(ctx, query.Q, &rel.ToUser, view) // `is_follow` always true
		if err != nil {
			return nil, err
		}
	}

	resp = &relation.RelationFollowListResponse{
		UserList: userList,
	}
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	r := query.FollowRelation
	u := query.User

	relList, err := r.WithContext(ctx).Preload(r.User).Where(r.ToUserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	view, err := u.WithContext(ctx).Where(u.ID.Eq(req.UserViewId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i], err = convert.UserConverterORM(ctx, query.Q, &rel.User, view) // `is_follow` uncertain
		if err != nil {
			return nil, err
		}
	}

	resp = &relation.RelationFollowerListResponse{
		UserList: userList,
	}
	return resp, nil
}
