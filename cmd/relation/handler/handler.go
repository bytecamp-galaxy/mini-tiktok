package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/relation"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	_, err = pack.QueryUser(ctx, req.ToUserId)
	if err != nil {
		return nil, err
	}

	// do action
	resp = &relation.RelationActionResponse{}
	switch req.ActionType {
	case 1:
		{
			e := doFollow(ctx, req.UserId, req.ToUserId)
			if e != nil {
				return nil, e
			}
			return resp, nil
		}
	case 2:
		{
			e := doUnFollow(ctx, req.UserId, req.ToUserId)
			if e != nil {
				return nil, e
			}
			return resp, nil
		}
	default:
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), "unknown action type")
	}
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	_, err = pack.QueryUser(ctx, req.UserViewId)
	if err != nil {
		return nil, err
	}

	r := query.FollowRelation
	relList, err := r.WithContext(ctx).Preload(r.ToUser).Where(r.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i], err = convert.UserConverterORM(ctx, query.Q, &rel.ToUser, req.UserViewId) // `is_follow` always true
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
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	_, err = pack.QueryUser(ctx, req.UserViewId)
	if err != nil {
		return nil, err
	}

	r := query.FollowRelation
	relList, err := r.WithContext(ctx).Preload(r.User).Where(r.ToUserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	userList := make([]*rpcmodel.User, len(relList))
	for i, rel := range relList {
		userList[i], err = convert.UserConverterORM(ctx, query.Q, &rel.User, req.UserViewId) // `is_follow` uncertain
		if err != nil {
			return nil, err
		}
	}

	resp = &relation.RelationFollowerListResponse{
		UserList: userList,
	}
	return resp, nil
}
