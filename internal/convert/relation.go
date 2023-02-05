package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
)

func isFollow(ctx context.Context, q *query.Query, userId int64, userViewId int64) (bool, error) {
	existed, err := redis.FollowKeyExist(ctx, userViewId)
	if err != nil {
		return false, err
	}

	if existed {
		// query redis
		return redis.FollowKeyContains(ctx, userViewId, userId)
	}

	// query db
	rs, _ := q.FollowRelation.WithContext(ctx).Where(q.FollowRelation.UserID.Eq(userViewId)).Find()
	ids := make([]int64, len(rs))
	for i, r := range rs {
		ids[i] = r.ToUserID
	}

	// load to redis
	_, err = redis.FollowKeyAdd(ctx, userViewId, ids)
	if err != nil {
		return false, err
	}

	// query redis
	return redis.FollowKeyContains(ctx, userViewId, userId)
}

func isFavorite(ctx context.Context, q *query.Query, videoId int64, userViewId int64) (bool, error) {
	existed, err := redis.FavouriteKeyExist(ctx, userViewId)
	if err != nil {
		return false, err
	}

	if existed {
		// query redis
		return redis.FavouriteKeyContains(ctx, userViewId, videoId)
	}

	rs, _ := q.FavoriteRelation.WithContext(ctx).Where(q.FavoriteRelation.UserID.Eq(userViewId)).Find()
	ids := make([]int64, len(rs))
	for i, r := range rs {
		ids[i] = r.VideoID
	}

	// load to redis
	_, err = redis.FavouriteKeyAdd(ctx, userViewId, ids)
	if err != nil {
		return false, err
	}

	// query redis
	return redis.FavouriteKeyContains(ctx, userViewId, videoId)
}
