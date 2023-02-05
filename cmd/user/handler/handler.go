package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// UserRegister implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	// hash password
	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrPasswordHash), err.Error())
	}

	// create user in db
	id := snowflake.Generate()
	err = query.User.WithContext(ctx).Create(&model.User{
		ID:       id,
		Username: req.Username,
		Password: hash,
	})
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// response to user server
	resp = &user.UserRegisterResponse{
		UserId: id,
	}
	return resp, nil
}

// UserLogin implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	// query user in db
	u := query.User
	data, err := query.User.WithContext(ctx).Where(u.Username.Eq(req.Username)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// check password
	if !utils.CheckPasswordHash(req.Password, data.Password) {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrPasswordIncorrect), "")
	}

	resp = &user.UserLoginResponse{
		UserId: data.ID,
	}
	return resp, nil
}

// UserQuery implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserQuery(ctx context.Context, req *user.UserQueryRequest) (resp *user.UserQueryResponse, err error) {
	// query user in db
	u, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(req.UserId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	view, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(req.UserViewId)).Take()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	resp = &user.UserQueryResponse{
		User: convert.UserConverterORM(ctx, query.Q, u, view),
	}
	return resp, nil
}
