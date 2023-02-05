package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
)

// VideoConverterAPI convert *rpcmodel.Video to *api.Video
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

// VideoConverterORM convert *model.Videos to *rpcmodel.Videos
func VideoConverterORM(ctx context.Context, q *query.Query, video *model.Video, view *model.User) *rpcmodel.Video {
	if video == nil {
		return nil
	}
	relFavorite := false
	relFollow := false
	if view != nil {
		relFavorite, _ = isFavorite(ctx, q, view.ID, video.ID)
	}
	author := video.Author // preload required
	if view != nil {
		relFollow, _ = isFollow(ctx, q, view.ID, author.ID)
	}
	u := &rpcmodel.User{
		Id:            author.ID,
		Name:          author.Username,
		FollowCount:   author.FollowingCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      relFollow,
	}
	res := &rpcmodel.Video{
		Id:            video.ID,
		Author:        u,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    relFavorite,
		Title:         video.Title,
	}
	return res
}
