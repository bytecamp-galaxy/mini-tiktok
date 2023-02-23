package handler

import (
	"context"
	"errors"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func doFavorite(ctx context.Context, uid int64, vid int64, auid int64) error {
	// NOTE: DO NOT USE `query.Q`
	// otherwise `favoriteList` reports "sql: transaction has already been committed or rolled back"

	// db transaction
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 添加点赞数据
		data := &model.FavoriteRelation{
			UserID:  uid,
			VideoID: vid,
		}
		err := tx.FavoriteRelation.WithContext(ctx).Create(data)
		if err != nil {
			return err
		}

		// update video FavoriteCount
		result, err := tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		// update user FavoriteCount
		result, err = tx.User.WithContext(ctx).
			Where(tx.User.ID.Eq(uid)).
			Update(tx.User.FavoriteCount, tx.User.FavoriteCount.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		// update author TotalFavorited
		result, err = tx.User.WithContext(ctx).
			Where(tx.User.ID.Eq(auid)).
			Update(tx.User.TotalFavorited, tx.User.TotalFavorited.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		return nil
	})

	if err != nil {
		return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// redis transaction
	err = func() error {
		// update redis favourite info if exists
		existed, err := redis.FavouriteKeyExists(ctx, uid)
		if err != nil {
			return err
		}
		if existed {
			err := redis.FavouriteKeyAdd(ctx, uid, vid)
			if err != nil {
				return err
			}
		}

		// update redis user info if exists
		exist, err := redis.UserInfoExists(ctx, uid)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, uid)
			if err != nil {
				return err
			}
			user.FavoriteCount += 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		exist, err = redis.UserInfoExists(ctx, auid)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, auid)
			if err != nil {
				return err
			}
			user.TotalFavorited += 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		return nil
	}()

	if err != nil {
		return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	return err
}

func doUnfavorite(ctx context.Context, uid int64, vid int64, auid int64) error {
	// db transaction
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 删除点赞数据
		r := tx.FavoriteRelation
		result, err := r.WithContext(ctx).Where(r.UserID.Eq(uid), r.VideoID.Eq(vid)).Delete()
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			return errors.New("nonexistent relation")
		}

		// update video FavoriteCount
		result, err = tx.Video.WithContext(ctx).
			Where(tx.Video.ID.Eq(vid)).
			Update(tx.Video.FavoriteCount, tx.Video.FavoriteCount.Sub(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		// update user FavoriteCount
		result, err = tx.User.WithContext(ctx).
			Where(tx.User.ID.Eq(uid)).
			Update(tx.User.FavoriteCount, tx.User.FavoriteCount.Sub(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		// update author TotalFavorited
		result, err = tx.User.WithContext(ctx).
			Where(tx.User.ID.Eq(auid)).
			Update(tx.User.TotalFavorited, tx.User.TotalFavorited.Sub(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		return nil
	})

	if err != nil {
		return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// redis transaction
	err = func() error {
		// update redis favourite info if exists
		existed, err := redis.FavouriteKeyExists(ctx, uid)
		if err != nil {
			return err
		}
		if existed {
			err := redis.FavouriteKeyRem(ctx, uid, vid)
			if err != nil {
				return err
			}
		}

		// update redis user info if exists
		exist, err := redis.UserInfoExists(ctx, uid)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, uid)
			if err != nil {
				return err
			}
			user.FavoriteCount -= 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		exist, err = redis.UserInfoExists(ctx, auid)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, auid)
			if err != nil {
				return err
			}
			user.TotalFavorited -= 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		return nil
	}()

	if err != nil {
		return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	return err
}
