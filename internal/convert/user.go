package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
)

// UserConverterAPI convert *rpcmodel.User to *api.User, can only be called by api servers
func UserConverterAPI(user *rpcmodel.User) *api.User {
	if user == nil {
		return nil
	}
	return &api.User{
		Id:              user.Id,
		Name:            user.Name,
		FollowCount:     &user.FollowCount,
		FollowerCount:   &user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          &user.Avatar,
		BackgroundImage: &user.BackgroundImage,
		Signature:       &user.Signature,
		TotalFavorited:  &user.TotalFavorited,
		WorkCount:       &user.WorkCount,
		FavoriteCount:   &user.FavoriteCount,
	}
}

// UserConverterORM convert *model.User to *rpcmodel.User, can only be called by rpc servers
func UserConverterORM(ctx context.Context, q *query.Query, user *model.User, userViewId int64) (res *rpcmodel.User, err error) {
	if user == nil {
		return nil, nil
	}
	relFollow := false
	if userViewId != redis.InvalidUserId {
		relFollow, err = pack.IsFollow(ctx, q, userViewId, user.ID)
		if err != nil {
			return nil, err
		}
	}
	return &rpcmodel.User{
		Id:              user.ID,
		Name:            user.Username,
		FollowCount:     user.FollowingCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        relFollow,
		Avatar:          "https://pixiv.cat/71001144.png",
		BackgroundImage: "https://pixiv.cat/105250474.png",
		Signature:       "Nothing but more and more nothingness.",
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}, nil
}
