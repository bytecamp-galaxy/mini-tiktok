package redis

import (
	"context"
)

var (
	followKeyFormat    = "ufo-%v"
	favouriteKeyFormat = "ufa-%v"
)

func FollowKeyExists(ctx context.Context, uid int64) (bool, error) {
	return Exists(ctx, followKeyFormat, uid)
}

func FollowKeyAdd(ctx context.Context, uid int64, id ...interface{}) error {
	return SetAdd(ctx, followKeyFormat, uid, id...)
}

func FollowKeyRem(ctx context.Context, uid int64, id ...interface{}) error {
	return SetRem(ctx, followKeyFormat, uid, id...)
}

func FollowKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return SetContains(ctx, followKeyFormat, uid, id)
}

func FavouriteKeyExists(ctx context.Context, uid int64) (bool, error) {
	return Exists(ctx, favouriteKeyFormat, uid)
}

func FavouriteKeyAdd(ctx context.Context, uid int64, id ...interface{}) error {
	return SetAdd(ctx, favouriteKeyFormat, uid, id...)
}

func FavouriteKeyRem(ctx context.Context, uid int64, id ...interface{}) error {
	return SetRem(ctx, favouriteKeyFormat, uid, id...)
}

func FavouriteKeyContains(ctx context.Context, uid int64, id int64) (bool, error) {
	return SetContains(ctx, favouriteKeyFormat, uid, id)
}
