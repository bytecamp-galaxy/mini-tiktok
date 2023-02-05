package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func doFavorite(ctx context.Context, uid int64, vid int64) error {
	// NOTE: DO NOT USE `query.Q`
	// otherwise `favoriteList` reports "sql: transaction has already been committed or rolled back"
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 1. 添加点赞数据
		u, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(uid)).Take()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		v, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).Take()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// TODO(vgalaxy): avoid insert video again
		err = tx.User.FavoriteVideos.WithContext(ctx).Model(u).Append(v)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Add(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
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
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		v, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).Take()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		err = tx.User.FavoriteVideos.WithContext(ctx).Model(u).Delete(v)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Sub(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		return nil
	})

	return err
}

func (s *FavoriteServiceImpl) favoriteList(ctx context.Context, req *favorite.FavoriteListRequest) ([]*rpcmodel.Video, error) {
	view, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(req.UserViewId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	u, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(req.UserId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	vs, err := query.User.FavoriteVideos.WithContext(ctx).Model(u).Find()
	// TODO(vgalaxy): preload author automatically
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	videos := make([]*rpcmodel.Video, len(vs))
	for i, v := range vs {
		author, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(v.AuthorID)).Take()
		if err != nil {
			return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		v.Author = *author // preload author manually
		videos[i] = convert.VideoConverterORM(ctx, query.Q, v, view)
	}

	return videos, nil
}
