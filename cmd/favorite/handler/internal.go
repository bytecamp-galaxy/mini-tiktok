package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func doFavorite(ctx context.Context, uid int64, vid int64) error {
	// NOTE: DO NOT USE `query.Q`
	// otherwise `favoriteList` reports "sql: transaction has already been committed or rolled back"
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 1. 添加点赞数据
		id := snowflake.Generate()
		err := tx.FavoriteRelation.WithContext(ctx).Create(&model.FavoriteRelation{
			ID:      id,
			UserID:  uid,
			VideoID: vid,
		})
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
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
		r := tx.FavoriteRelation
		result, err := r.WithContext(ctx).Where(r.UserID.Eq(uid), r.VideoID.Eq(vid)).Delete()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected == 0 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "nonexistent relation")
		}

		// 2.改变 video 表中的 FavoriteCount
		_, err = tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
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

	fr := query.FavoriteRelation
	rs, err := fr.WithContext(ctx).Preload(fr.Video).Preload(fr.Video.Author).Where(fr.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	videos := make([]*rpcmodel.Video, len(rs))
	for i, r := range rs {
		videos[i] = convert.VideoConverterORM(ctx, query.Q, &r.Video, view)
	}

	return videos, nil
}
