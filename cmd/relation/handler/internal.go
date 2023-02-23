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

func doFollow(ctx context.Context, fromId int64, toId int64) error {
	// db transaction
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 添加关注数据
		data := &model.FollowRelation{
			UserID:   fromId,
			ToUserID: toId,
		}
		err := tx.FollowRelation.WithContext(ctx).Create(data)
		if err != nil {
			return err
		}

		// 修改 FollowingCount 和 FollowerCount
		u := tx.User
		result, err := u.WithContext(ctx).Where(u.ID.Eq(fromId)).Update(u.FollowingCount, u.FollowingCount.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		result, err = u.WithContext(ctx).Where(u.ID.Eq(toId)).Update(u.FollowerCount, u.FollowerCount.Add(1))
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
		// update redis follow info if exists
		exist, err := redis.FollowKeyExists(ctx, fromId)
		if err != nil {
			return err
		}
		if exist {
			err = redis.FollowKeyAdd(ctx, fromId, toId)
			if err != nil {
				return err
			}
		}

		// update redis user info if exists
		exist, err = redis.UserInfoExists(ctx, fromId)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, fromId)
			if err != nil {
				return err
			}
			user.FollowingCount += 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		exist, err = redis.UserInfoExists(ctx, toId)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, toId)
			if err != nil {
				return err
			}
			user.FollowerCount += 1
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

	return nil
}

func doUnFollow(ctx context.Context, fromId int64, toId int64) error {
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 删除关注数据
		r := tx.FollowRelation
		result, err := r.WithContext(ctx).Where(r.UserID.Eq(fromId), r.ToUserID.Eq(toId)).Delete()
		if err != nil {
			return err
		}
		if result.RowsAffected == 0 {
			return errors.New("nonexistent relation")
		}

		// 修改 FollowingCount 和 FollowerCount
		u := tx.User
		result, err = u.WithContext(ctx).Where(u.ID.Eq(fromId)).Update(u.FollowingCount, u.FollowingCount.Sub(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		result, err = u.WithContext(ctx).Where(u.ID.Eq(toId)).Update(u.FollowerCount, u.FollowerCount.Sub(1))
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
		// update redis follow info if exists
		exist, err := redis.FollowKeyExists(ctx, fromId)
		if err != nil {
			return err
		}
		if exist {
			err = redis.FollowKeyRem(ctx, fromId, toId)
			if err != nil {
				return err
			}
		}

		// update redis user info if exists
		exist, err = redis.UserInfoExists(ctx, fromId)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, fromId)
			if err != nil {
				return err
			}
			user.FollowingCount -= 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		exist, err = redis.UserInfoExists(ctx, toId)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, toId)
			if err != nil {
				return err
			}
			user.FollowerCount -= 1
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

	return nil
}
