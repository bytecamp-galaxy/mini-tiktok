package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
)

// VideoConverterAPI convert *rpcmodel.Video to *api.Video, can only be called by api servers
func VideoConverterAPI(video *rpcmodel.Video) *api.Video {
	if video == nil {
		return nil
	}
	author := video.Author
	u := &api.User{
		Id:            author.Id,
		Name:          author.Name,
		FollowCount:   &author.FollowCount,
		FollowerCount: &author.FollowerCount,
		IsFollow:      author.IsFollow,
	}
	res := &api.Video{
		Id:            video.Id,
		Author:        u,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    video.IsFavorite,
		Title:         video.Title,
	}
	return res
}

// VideoConverterORM convert *model.Videos to *rpcmodel.Videos, can only be called by rpc servers
func VideoConverterORM(ctx context.Context, q *query.Query, video *model.Video, view *model.User) (res *rpcmodel.Video, err error) {
	if video == nil {
		return nil, nil
	}

	relFavorite := false
	relFollow := false
	author := video.Author // preload required
	if view != nil {
		relFavorite, err = isFavorite(ctx, q, view.ID, video.ID)
		if err != nil {
			return nil, err
		}
		relFollow, err = isFollow(ctx, q, view.ID, author.ID)
		if err != nil {
			return nil, err
		}
	}

	u := &rpcmodel.User{
		Id:            author.ID,
		Name:          author.Username,
		FollowCount:   author.FollowingCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      relFollow,
	}
	res = &rpcmodel.Video{
		Id:            video.ID,
		Author:        u,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    relFavorite,
		Title:         video.Title,
	}
	return res, nil
}
