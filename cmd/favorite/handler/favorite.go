package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"gorm.io/gorm"
)

func Favorite(ctx context.Context, uid int64, vid int64) error {
	err := mysql.DB.Transaction(func(tx *gorm.DB) error {
		user := new(model.User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video := new(model.Video)
		if err := tx.WithContext(ctx).First(video, vid).Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).Model(&user).Association("FavoriteVideos").Append(video); err != nil {
			return err
		}
		// 2.改变 video 表中的 FavoriteCount
		res := tx.Model(video).Update("FavoriteCount", gorm.Expr("FavoriteCount + ?", 1))
		if res.Error != nil {
			return res.Error
		}

		return nil
	})
	return err
}

func DisFavorite(ctx context.Context, uid int64, vid int64) error {
	err := mysql.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// 1. 删除点赞数据
		user := new(model.User)
		if err := tx.WithContext(ctx).First(user, uid).Error; err != nil {
			return err
		}

		video, err := GetFavoriteRelation(ctx, uid, vid)
		if err != nil {
			return err
		}

		err = tx.Unscoped().WithContext(ctx).Model(&user).Association("FavoriteVideos").Delete(video)
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 FavoriteCount
		res := tx.Model(video).Update("FavoriteCount", gorm.Expr("FavoriteCount - ?", 1))
		if res.Error != nil {
			return res.Error
		}

		return nil
	})
	return err
}

// GetFavoriteRelation get favorite video info
func GetFavoriteRelation(ctx context.Context, uid int64, vid int64) (*model.Video, error) {
	user := new(model.User)
	if err := mysql.DB.WithContext(ctx).First(user, uid).Error; err != nil {
		return nil, err
	}

	video := new(model.Video)

	if err := mysql.DB.WithContext(ctx).Model(&user).Association("FavoriteVideos").Find(&video, vid); err != nil {
		return nil, err
	}
	return video, nil
}

// FavoriteList returns a list of Favorite videos.
func (s *FavoriteServiceImpl) _FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*rpcmodel.Video, error) {
	var Favoritevideos []model.Video
	videos, err := FavoriteVideos(ctx, Favoritevideos, &req.UserId)
	return videos, err
}

func FavoriteVideos(ctx context.Context, vs []model.Video, uid *int64) ([]*rpcmodel.Video, error) {
	videos := make([]*model.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)

	return packVideos, err
}

func Videos(ctx context.Context, vs []*model.Video, fromID *int64) ([]*rpcmodel.Video, error) {
	respVideos := make([]*rpcmodel.Video, len(vs))
	for i, video := range vs {
		author := video.Author
		u := &rpcmodel.User{
			Id:            author.ID,
			Name:          author.Username,
			FollowCount:   author.FollowingCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      false, // TODO
		}
		respVideos[i] = &rpcmodel.Video{
			Id:            video.ID,
			Author:        u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false, // TODO: 如果用户登录状态下刷视频，如何高效的获取这些用户对刷到的视频的点赞信息？
			Title:         video.Title,
		}
	}
	return respVideos, nil
}
