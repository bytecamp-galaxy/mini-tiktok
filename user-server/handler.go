package main

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	user "github.com/bytecamp-galaxy/mini-tiktok/user-server/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		panic(err)
	}

	err = query.User.WithContext(ctx).Create(&model.User{
		UserName: req.Username,
		Password: hash,
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

	if !utils.CheckPasswordHash(req.Password, data.Password) {
		panic("incorrect password")
	}

	resp = &user.UserLoginResponse{
		StatusCode: 0,
		StatusMsg:  "success",
		UserId:     int64(data.ID),
	}
	return resp, nil
}

// UserQuery implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserQuery(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	// TODO: Your code here...
	return
}
