package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/relation"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
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
				id := snowflake.Generate()
				err := tx.Relation.WithContext(ctx).Create(&model.Relation{
					ID:       id,
					UserID:   req.UserId,
					ToUserID: req.ToUserId,
				})
				if err != nil {
					return err
				}

				u := tx.User
				_, err = u.WithContext(ctx).Where(u.ID.Eq(req.UserId)).Update(u.FollowingCount, u.FollowingCount.Add(1))
				if err != nil {
					return err
				}
				_, err = u.WithContext(ctx).Where(u.ID.Eq(req.ToUserId)).Update(u.FollowerCount, u.FollowerCount.Add(1))
				if err != nil {
					return err
				}
			}
		case 2:
			{
				r := tx.Relation
				_, err := r.WithContext(ctx).Where(r.UserID.Eq(req.UserId), r.ToUserID.Eq(req.ToUserId)).Delete()
				if err != nil {
					return err
				}

				u := tx.User
				_, err = u.WithContext(ctx).Where(u.ID.Eq(req.UserId)).Update(u.FollowingCount, u.FollowingCount.Sub(1))
				if err != nil {
					return err
				}
				_, err = u.WithContext(ctx).Where(u.ID.Eq(req.ToUserId)).Update(u.FollowerCount, u.FollowerCount.Sub(1))
				if err != nil {
					return err
				}
			}
		default:
			return kerrors.NewBizStatusError(int32(errno.ErrUnknown), "request argument violates convention")
		}
		return nil
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	return &relation.RelationActionResponse{}, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	r := query.Relation
	relList, err := r.WithContext(ctx).Preload(r.ToUser).Where(r.UserID.Eq(req.UserId)).Find()

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i] = convert.UserConverterORM(&rel.ToUser) // `is_follow` always true
	}

	resp = &relation.RelationFollowListResponse{
		UserList: userList,
	}
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	r := query.Relation
	relList, err := r.WithContext(ctx).Preload(r.User).Where(r.UserID.Eq(req.UserId)).Find()

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i] = convert.UserConverterORM(&rel.User) // `is_follow` uncertain
	}

	resp = &relation.RelationFollowerListResponse{
		UserList: userList,
	}
	return resp, nil
}
