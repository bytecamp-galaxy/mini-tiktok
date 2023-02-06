package redis

import (
	"context"
)

var (
	followKeyFormat    = "follow-%v"
	favouriteKeyFormat = "favourite-%v"
)

func FollowKeyExist(ctx context.Context, uid int64) (bool, error) {
	return Exist(followKeyFormat, ctx, uid)
}

func FollowKeyAdd(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return SetAdd(followKeyFormat, ctx, uid, id...)
}

func FollowKeyRem(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return SetRem(followKeyFormat, ctx, uid, id...)
}

func FollowKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return SetContains(followKeyFormat, ctx, uid, id)
}

func FavouriteKeyExist(ctx context.Context, uid int64) (bool, error) {
	return Exist(favouriteKeyFormat, ctx, uid)
}

func FavouriteKeyAdd(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return SetAdd(favouriteKeyFormat, ctx, uid, id...)
}

func FavouriteKeyRem(ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	return SetRem(favouriteKeyFormat, ctx, uid, id...)
}

func FavouriteKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return SetContains(favouriteKeyFormat, ctx, uid, id)
}
