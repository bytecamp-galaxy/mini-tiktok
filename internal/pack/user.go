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
	// query user id in redis bloom filter
	exist, err := redis.UserIdExistBF(ctx, uid)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}
	if !exist {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidUser), "")
	}

	// query user info in redis
	exist, err = redis.UserInfoExists(ctx, uid)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	var u *model.User
	if exist {
		u, err = redis.UserInfoGet(ctx, uid)
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
		err := redis.UserInfoSet(ctx, u)
		if err != nil {
			return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
	}
	return u, nil
}
