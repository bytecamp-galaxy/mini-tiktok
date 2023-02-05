package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
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
	isFavorite := false
	isFollow := false
	if view != nil {
		count, _ := q.FavoriteRelation.WithContext(ctx).
			Where(q.FavoriteRelation.UserID.Eq(view.ID), q.FavoriteRelation.VideoID.Eq(video.ID)).
			Count()
		if count != 0 {
			isFavorite = true
		}
	}
	author := video.Author // preload required
	if view != nil {
		count, _ := q.FollowRelation.WithContext(ctx).
			Where(q.FollowRelation.UserID.Eq(view.ID), q.FollowRelation.ToUserID.Eq(author.ID)).
			Count()
		if count != 0 {
			isFollow = true
		}
	}
	u := &rpcmodel.User{
		Id:            author.ID,
		Name:          author.Username,
		FollowCount:   author.FollowingCount,
		FollowerCount: author.FollowerCount,
		IsFollow:      isFollow,
	}
	res := &rpcmodel.Video{
		Id:            video.ID,
		Author:        u,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    isFavorite,
		Title:         video.Title,
	}
	return res
}
