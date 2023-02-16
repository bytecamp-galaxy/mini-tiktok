package redis

import (
	"context"
	"fmt"
	"github.com/marmotedu/errors"
)

/*==================================================================
                           Common
====================================================================*/

func Exists(ctx context.Context, format string, uid int64) (bool, error) {
	key := fmt.Sprintf(format, uid)
	existed, err := r.Exists(ctx, key).Result()
	return existed == 1, err
}

func Del(ctx context.Context, format string, uid int64) error {
	key := fmt.Sprintf(format, uid)
	_, err := r.Del(ctx, key).Result() // skip checking result equals one
	if err != nil {
		return err
	}
	return nil
}

/*==================================================================
                             Set
====================================================================*/

func SetAdd(ctx context.Context, format string, uid int64, id ...interface{}) error {
	key := fmt.Sprintf(format, uid)
	result, err := r.SAdd(ctx, key, id...).Result()
	if err != nil {
		return err
	}
	if result != int64(len(id)) {
		return errors.New("redis sadd error")
	}
	return nil
}

func SetRem(ctx context.Context, format string, uid int64, id ...interface{}) error {
	key := fmt.Sprintf(format, uid)
	result, err := r.SRem(ctx, key, id...).Result()
	if err != nil {
		return err
	}
	if result != int64(len(id)) {
		return errors.New("redis srem error")
	}
	return nil
}

func SetContains(ctx context.Context, format string, uid int64, id int64) (bool, error) {
	key := fmt.Sprintf(format, uid)
	return r.SIsMember(ctx, key, id).Result()
}

/*==================================================================
                         Bloom Filter
====================================================================*/

func BFInit(ctx context.Context, key string) error {
	result, err := r.Do(ctx, "bf.reserve", key, 0.001 /* error rate */, 50000 /* initial size */).Result()
	if err != nil {
		return err
	}
	if result.(string) != "OK" {
		return errors.New("redis bf.reserve error")
	}
	return nil
}

func BFAdd(ctx context.Context, key string, value interface{}) error {
	result, err := r.Do(ctx, "bf.add", key, value).Result()
	if err != nil {
		return err
	}
	if result.(int64) != 1 {
		return errors.New("redis bf.add error")
	}
	return nil
}

func BFExists(ctx context.Context, key string, value interface{}) (bool, error) {
	result, err := r.Do(ctx, "bf.exists", key, value).Result()
	if err != nil {
		return false, err
	}
	return result.(int64) == 1, nil
}
