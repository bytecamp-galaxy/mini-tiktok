package pack

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// QueryUser can only be called by rpc servers
func QueryUser(ctx context.Context, uid int64) (*model.User, error) {
	var u *model.User
	exist, err := redis.UserKeyExist(ctx, uid)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	if exist {
		u, err = redis.UserKeyGet(ctx, uid)
		if err != nil {
			return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
	} else {
		// query user in db
		u, _ = query.User.WithContext(ctx).Where(query.User.ID.Eq(uid)).Take()
		if u == nil {
			return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidUser), "")
		}

		// load to redis
		err := redis.UserKeySet(ctx, u)
		if err != nil {
			return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
	}
	return u, nil
}

func DeleteUserFromRedisIfExist(ctx context.Context, uid int64) error {
	existed, err := redis.UserKeyExist(ctx, uid)
	if err != nil {
		return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}
	if existed {
		count, err := redis.UserKeyDel(ctx, uid)
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if count != 1 {
			return kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis del error")
		}
	}
	return nil
}
