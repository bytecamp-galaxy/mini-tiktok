package redis

import (
	"context"
)

var (
	followKeyFormat    = "follow-%v"
	favouriteKeyFormat = "favourite-%v"
)

func FollowKeyExist(ctx context.Context, uid int64) (bool, error) {
	return setExist(followKeyFormat, ctx, uid)
}

func FollowKeyAdd(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return setAdd(followKeyFormat, ctx, uid, id...)
}

func FollowKeyRem(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return setRem(followKeyFormat, ctx, uid, id...)
}

func FollowKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return setContains(followKeyFormat, ctx, uid, id)
}

func FavouriteKeyExist(ctx context.Context, uid int64) (bool, error) {
	return setExist(favouriteKeyFormat, ctx, uid)
}

func FavouriteKeyAdd(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return setAdd(favouriteKeyFormat, ctx, uid, id...)
}

func FavouriteKeyRem(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return setRem(favouriteKeyFormat, ctx, uid, id...)
}

func FavouriteKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return setContains(favouriteKeyFormat, ctx, uid, id)
}
