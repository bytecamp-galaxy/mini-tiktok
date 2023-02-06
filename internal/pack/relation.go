package pack

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// IsFollow user (view) follow user
func IsFollow(ctx context.Context, q *query.Query, userViewId int64, userId int64) (bool, error) {
	existed, err := redis.FollowKeyExist(ctx, userViewId)
	if err != nil {
		return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	if !existed {
		// query db
		rs, err := q.FollowRelation.WithContext(ctx).Where(q.FollowRelation.UserID.Eq(userViewId)).Find()
		if err != nil {
			return false, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		ids := make([]interface{}, len(rs)+1)
		ids[0] = redis.InvalidUserId
		for i, r := range rs {
			ids[i] = r.ToUserID
		}

		// load to redis
		count, err := redis.FollowKeyAdd(ctx, userViewId, ids...)
		if err != nil {
			return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if count != int64(len(rs)+1) {
			return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis sadd error")
		}
	}

	// query redis
	res, err := redis.FollowKeyContains(ctx, userViewId, userId)
	if err != nil {
		return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}
	return res, nil
}

// IsFavorite user (view) favorite video
func IsFavorite(ctx context.Context, q *query.Query, userViewId int64, videoId int64) (bool, error) {
	existed, err := redis.FavouriteKeyExist(ctx, userViewId)
	if err != nil {
		return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	if !existed {
		// query db
		rs, err := q.FavoriteRelation.WithContext(ctx).Where(q.FavoriteRelation.UserID.Eq(userViewId)).Find()
		if err != nil {
			return false, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}
		ids := make([]interface{}, len(rs)+1)
		ids[0] = redis.InvalidUserId
		for i, r := range rs {
			ids[i] = r.VideoID
		}

		// load to redis
		count, err := redis.FavouriteKeyAdd(ctx, userViewId, ids...)
		if err != nil {
			return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
		}
		if count != int64(len(rs)+1) {
			return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), "redis sadd error")
		}
	}

	// query redis
	res, err := redis.FavouriteKeyContains(ctx, userViewId, videoId)
	if err != nil {
		return false, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}
	return res, nil
}
