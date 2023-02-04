package convert

import (
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
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
func UserConverterORM(user *model.User) *rpcmodel.User {
	if user == nil {
		return nil
	}
	return &rpcmodel.User{
		Id:            user.ID,
		Name:          user.Username,
		FollowCount:   user.FollowingCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false, // TODO(vgalaxy)
	}
}
