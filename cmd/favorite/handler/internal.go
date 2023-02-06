package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func doFavorite(ctx context.Context, uid int64, vid int64) error {
	// NOTE: DO NOT USE `query.Q`
	// otherwise `favoriteList` reports "sql: transaction has already been committed or rolled back"
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 添加点赞数据
		id := snowflake.Generate()
		err := tx.FavoriteRelation.WithContext(ctx).Create(&model.FavoriteRelation{
			ID:      id,
			UserID:  uid,
			VideoID: vid,
		})
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// 改变 video 表中的 FavoriteCount
		result, err := tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Add(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		// update redis favourite info if exists
		existed, err := redis.FavouriteKeyExist(ctx, uid)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if existed {
			count, err := redis.FavouriteKeyAdd(ctx, uid, vid)
			if err != nil {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
			}
			if count != 1 {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis sadd error")
			}
		}

		return nil
	})

	return err
}

func doUnfavorite(ctx context.Context, uid int64, vid int64) error {
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 删除点赞数据
		r := tx.FavoriteRelation
		result, err := r.WithContext(ctx).Where(r.UserID.Eq(uid), r.VideoID.Eq(vid)).Delete()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected == 0 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "nonexistent relation")
		}

		// 改变 video 表中的 FavoriteCount
		result, err = tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Sub(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		// update redis favourite info if exists
		existed, err := redis.FavouriteKeyExist(ctx, uid)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if existed {
			count, err := redis.FavouriteKeyRem(ctx, uid, vid)
			if err != nil {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
			}
			if count != 1 {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis srem error")
			}
		}

		return nil
	})

	return err
}
