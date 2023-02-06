package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
)

// UserConverterAPI convert *rpcmodel.User to *api.User, can only be called by api servers
func UserConverterAPI(user *rpcmodel.User) *api.User {
	if user == nil {
		return nil
	}
	return &api.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   &user.FollowCount,
		FollowerCount: &user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

// UserConverterORM convert *model.User to *rpcmodel.User, can only be called by rpc servers
func UserConverterORM(ctx context.Context, q *query.Query, user *model.User, view *model.User) (res *rpcmodel.User, err error) {
	if user == nil {
		return nil, nil
	}
	relFollow := false
	if view != nil {
		relFollow, err = isFollow(ctx, q, view.ID, user.ID)
		if err != nil {
			return nil, err
		}
	}
	return &rpcmodel.User{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   user.FollowingCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      relFollow,
	}, nil
}
