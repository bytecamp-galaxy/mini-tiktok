package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/cloudwego/kitex/pkg/klog"
)

func doFavorite(ctx context.Context, uid int64, vid int64) error {
	// NOTE: DO NOT USE `query.Q`
	// otherwise `favoriteList` reports "sql: transaction has already been committed or rolled back"
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 1. 添加点赞数据
		u, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(uid)).Take()
		if err != nil {
			return err
		}

		v, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).Take()
		if err != nil {
			return err
		}

		// TODO(vgalaxy): avoid insert video again
		err = tx.User.FavoriteVideos.WithContext(ctx).Model(u).Append(v)
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Add(1))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func doUnfavorite(ctx context.Context, uid int64, vid int64) error {
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 1. 删除点赞数据
		u, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(uid)).Take()
		if err != nil {
			return err
		}

		v, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).Take()
		if err != nil {
			return err
		}

		err = tx.User.FavoriteVideos.WithContext(ctx).Model(u).Delete(v)
		if err != nil {
			return err
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Sub(1))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (s *FavoriteServiceImpl) favoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*rpcmodel.Video, error) {
	q := query.Q
	u, err := q.User.WithContext(ctx).Where(q.User.ID.Eq(req.UserId)).Take()
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	vs, err := q.User.FavoriteVideos.WithContext(ctx).Model(u).Find()
	if err != nil {
		klog.CtxErrorf(ctx, err.Error())
		return nil, err
	}

	videos := make([]*rpcmodel.Video, len(vs))
	for i, v := range vs {
		// TODO(vgalaxy): optimize join
		u, err := q.User.WithContext(ctx).Where(q.User.ID.Eq(v.AuthorID)).Take()
		if err != nil {
			klog.CtxErrorf(ctx, err.Error())
			return nil, err
		}
		author := &rpcmodel.User{
			Id:            u.ID,
			Name:          u.Username,
			FollowCount:   u.FollowingCount,
			FollowerCount: u.FollowerCount,
			IsFollow:      false, // TODO
		}
		videos[i] = &rpcmodel.Video{
			Id:            v.ID,
			Author:        author,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    true,
			Title:         v.Title,
		}
	}

	return videos, err
}
