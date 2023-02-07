package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	InvalidUserId int64 = -1
)

const (
	userInfoKeyFormat string = "u-%v"
	userIdExistKey    string = "uid"
	userNameExistKey  string = "uname"
)

// LoadUserFromDBToRedis only called by api service
func LoadUserFromDBToRedis(ctx context.Context) error {
	// init bloom filter
	err := UserIdInitBF(ctx)
	if err != nil {
		return err
	}
	err = UserNameInitBF(ctx)
	if err != nil {
		return err
	}
	// query db
	us, err := query.User.WithContext(ctx).Find()
	if err != nil {
		return err
	}
	// foreach
	for _, u := range us {
		// load user id
		err = UserIdAddBF(ctx, u.ID)
		if err != nil {
			return err
		}
		// load user name
		err = UserNameAddBF(ctx, u.Username)
		if err != nil {
			return err
		}
		// load user info
		err := UserInfoSet(ctx, u)
		if err != nil {
			return err
		}
	}
	hlog.CtxInfof(ctx, "load %v user(s) from db to redis successfully", len(us))
	return nil
}

/*==================================================================
                          User Id
====================================================================*/

func UserIdInitBF(ctx context.Context) error {
	return BFInit(ctx, userIdExistKey)
}

func UserIdExistBF(ctx context.Context, uid int64) (bool, error) {
	return BFExists(ctx, userIdExistKey, uid)
}

func UserIdAddBF(ctx context.Context, uid int64) error {
	return BFAdd(ctx, userIdExistKey, uid)
}

/*==================================================================
                          User Name
====================================================================*/

func UserNameInitBF(ctx context.Context) error {
	return BFInit(ctx, userNameExistKey)
}

func UserNameExistBF(ctx context.Context, uname string) (bool, error) {
	return BFExists(ctx, userNameExistKey, uname)
}

func UserNameAddBF(ctx context.Context, uname string) error {
	return BFAdd(ctx, userNameExistKey, uname)
}

/*==================================================================
                          User Info
====================================================================*/

func UserInfoExists(ctx context.Context, uid int64) (bool, error) {
	return Exists(ctx, userInfoKeyFormat, uid)
}

func UserInfoSet(ctx context.Context, user *model.User) error {
	k := fmt.Sprintf(userInfoKeyFormat, user.ID)
	v, err := json.Marshal(*user)
	if err != nil {
		return err
	}
	res, err := r.Set(ctx, k, v, 0).Result()
	if err != nil {
		return err
	}
	if res != "OK" {
		return errors.New("redis set error")
	}
	return nil
}

func UserInfoDel(ctx context.Context, uid int64) error {
	k := fmt.Sprintf(userInfoKeyFormat, uid)
	res, err := r.Del(ctx, k).Result()
	if err != nil {
		return err
	}
	if res != 1 {
		return errors.New("redis del error")
	}
	return nil
}

func UserInfoGet(ctx context.Context, uid int64) (*model.User, error) {
	k := fmt.Sprintf(userInfoKeyFormat, uid)
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
