package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
)

// UserConverterAPI convert *rpcmodel.User to *api.User
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

// UserConverterORM convert *model.User to *rpcmodel.User
func UserConverterORM(ctx context.Context, q *query.Query, user *model.User, view *model.User) *rpcmodel.User {
	if user == nil {
		return nil
	}
	isFollow := false
	if view != nil {
		count, _ := q.Relation.WithContext(ctx).Where(q.Relation.UserID.Eq(view.ID), q.Relation.ToUserID.Eq(user.ID)).Count()
		if count != 0 {
			isFollow = true
		}
	}
	return &rpcmodel.User{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   user.FollowingCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
}
