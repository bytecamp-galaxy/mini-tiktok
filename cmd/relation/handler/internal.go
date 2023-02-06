package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func doFollow(ctx context.Context, fromId int64, toId int64) error {
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 添加关注数据
		id := snowflake.Generate()
		err := tx.FollowRelation.WithContext(ctx).Create(&model.FollowRelation{
			ID:       id,
			UserID:   fromId,
			ToUserID: toId,
		})
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		// 修改 FollowingCount 和 FollowerCount
		u := tx.User
		result, err := u.WithContext(ctx).Where(u.ID.Eq(fromId)).Update(u.FollowingCount, u.FollowingCount.Add(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		result, err = u.WithContext(ctx).Where(u.ID.Eq(toId)).Update(u.FollowerCount, u.FollowerCount.Add(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		// update redis follow info if exists
		exist, err := redis.FollowKeyExist(ctx, fromId)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if exist {
			count, err := redis.FollowKeyAdd(ctx, fromId, toId)
			if err != nil {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
			}
			if count != 1 {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis sadd error")
			}
		}

		// delete redis user info if exists
		err = pack.DeleteUserFromRedisIfExist(ctx, fromId)
		if err != nil {
			return err
		}

		err = pack.DeleteUserFromRedisIfExist(ctx, toId)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func doUnFollow(ctx context.Context, fromId int64, toId int64) error {
	q := query.Use(mysql.DB)
	err := q.Transaction(func(tx *query.Query) error {
		// 删除关注数据
		r := tx.FollowRelation
		result, err := r.WithContext(ctx).Where(r.UserID.Eq(fromId), r.ToUserID.Eq(toId)).Delete()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected == 0 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "nonexistent relation")
		}

		// 修改 FollowingCount 和 FollowerCount
		u := tx.User
		result, err = u.WithContext(ctx).Where(u.ID.Eq(fromId)).Update(u.FollowingCount, u.FollowingCount.Sub(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		result, err = u.WithContext(ctx).Where(u.ID.Eq(toId)).Update(u.FollowerCount, u.FollowerCount.Sub(1))
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		if result.RowsAffected != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
		}

		// update redis follow info if exists
		exist, err := redis.FollowKeyExist(ctx, fromId)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if exist {
			count, err := redis.FollowKeyRem(ctx, fromId, toId)
			if err != nil {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
			}
			if count != 1 {
				return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis srem error")
			}
		}

		// delete redis user info if exists
		err = pack.DeleteUserFromRedisIfExist(ctx, fromId)
		if err != nil {
			return err
		}

		err = pack.DeleteUserFromRedisIfExist(ctx, toId)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
