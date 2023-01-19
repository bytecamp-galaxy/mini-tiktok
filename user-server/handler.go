package main

import (
	"context"
	"mini-tiktok-v2/pkg/dal/model"
	"mini-tiktok-v2/pkg/dal/query"
	"mini-tiktok-v2/pkg/utils"
	user "mini-tiktok-v2/user-server/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	err = query.User.WithContext(ctx).Create(&model.User{
		UserName: req.Username,
		Password: utils.MD5(req.Password),
	})
	if err != nil {
		panic(err)
	}

	q := query.Q
	t := q.User

	data, err := query.User.WithContext(ctx).Where(t.UserName.Eq(req.Username)).Take()
	if err != nil {
		panic(err)
	}

	resp = &user.UserRegisterResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(data.ID),
		Token:      "",
	}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	q := query.Q
	t := q.User

	data, err := query.User.WithContext(ctx).Where(t.UserName.Eq(req.Username)).Take()
	if err != nil {
		panic(err)
	}

	if data.Password != utils.MD5(req.Password) {
		panic("incorrect password")
	}

	resp = &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(data.ID),
		Token:      "",
	}
	return resp, nil
}
