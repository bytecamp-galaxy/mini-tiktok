package redis

import (
	"context"
	"fmt"
)

func setExist(format string, ctx context.Context, uid int64) (bool, error) {
	key := fmt.Sprintf(format, uid)
	existed, err := r.Exists(ctx, key).Result()
	return existed == 1, err
}

func setAdd(format string, ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	key := fmt.Sprintf(format, uid)
	return r.SAdd(ctx, key, id...).Result()
}

func setRem(format string, ctx context.Context, uid int64, id ...interface{}) (int64, error) {
	key := fmt.Sprintf(format, uid)
	return r.SRem(ctx, key, id...).Result()
}

func setContains(format string, ctx context.Context, uid int64, id int64) (bool, error) {
	key := fmt.Sprintf(format, uid)
	return r.SIsMember(ctx, key, id).Result()
}
