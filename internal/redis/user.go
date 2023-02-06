package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	InvalidUserId int64 = -1
)

const (
	userKeyFormat = "user-%v"
)

// LoadUserFromDBToRedis only called by api service
func LoadUserFromDBToRedis(ctx context.Context) error {
	us, err := query.User.WithContext(ctx).Find()
	if err != nil {
		return err
	}
	for _, u := range us {
		err := UserKeySet(ctx, u)
		if err != nil {
			return err
		}
	}
	hlog.CtxInfof(ctx, "load %v user(s) from db to redis successfully", len(us))
	return nil
}

func UserKeyExist(ctx context.Context, uid int64) (bool, error) {
	return Exist(userKeyFormat, ctx, uid)
}

func UserKeySet(ctx context.Context, user *model.User) error {
	k := fmt.Sprintf(userKeyFormat, user.ID)
	v, err := json.Marshal(*user)
	if err != nil {
		return err
	}
	_, err = r.Set(ctx, k, v, 0).Result()
	if err != nil {
		return err
	}
	return nil
}

func UserKeyDel(ctx context.Context, uid int64) (int64, error) {
	k := fmt.Sprintf(userKeyFormat, uid)
	return r.Del(ctx, k).Result()
}

func UserKeyGet(ctx context.Context, uid int64) (*model.User, error) {
	k := fmt.Sprintf(userKeyFormat, uid)
	res, err := r.Get(ctx, k).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
