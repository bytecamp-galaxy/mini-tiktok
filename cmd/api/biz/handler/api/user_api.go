// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/marmotedu/errors"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// validate password
	err = utils.ValidatePassword(req.Password)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrPasswordInvalid, err.Error()))
		return
	}

	// set up connection with user server
	cli, err := rpc.InitUserClient()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	// call rpc service
	reqRpc := &user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := (*cli).UserRegister(ctx, reqRpc)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			e := errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage())
			pack.Error(c, errors.WrapC(e, errno.ErrRPCProcess, ""))
			return
		} else {
			// assume
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	// generate token
	token, _, err := jwt.Middleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrTokenGeneration, err.Error()))
		return
	}

	// response to client
	resp := &api.UserRegisterResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		UserId:     respRpc.UserId,
		Token:      token,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/user/login/ [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// set up connection with user server
	cli, err := rpc.InitUserClient()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	// call rpc service
	reqRpc := &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := (*cli).UserLogin(ctx, reqRpc)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			e := errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage())
			pack.Error(c, errors.WrapC(e, errno.ErrRPCProcess, ""))
			return
		} else {
			// assume
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	// generate token
	token, _, err := jwt.Middleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrTokenGeneration, err.Error()))
		return
	}

	// response to client
	resp := &api.UserLoginResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		UserId:     respRpc.UserId,
		Token:      token,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserQuery .
// @router /douyin/user/ [GET]
func UserQuery(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// fetch user id from token
	id, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, pack.BrokenInvariantStatusMessage))
		return
	}

	// check user id
	if id != req.UserId {
		pack.Error(c, errors.WithCode(errno.ErrTokenInvalid, "inconsistent user id"))
		return
	}

	// set up connection with user server
	cli, err := rpc.InitUserClient()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	// call rpc service
	reqRpc := &user.UserQueryRequest{
		UserId: req.UserId,
	}

	respRpc, err := (*cli).UserQuery(ctx, reqRpc)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			e := errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage())
			pack.Error(c, errors.WrapC(e, errno.ErrRPCProcess, ""))
			return
		} else {
			// assume
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	// response to client
	resp := &api.UserQueryResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		User: &api.User{
			Id:            respRpc.User.Id,
			Name:          respRpc.User.Name,
			FollowCount:   utils.Int64(respRpc.User.FollowerCount),
			FollowerCount: utils.Int64(respRpc.User.FollowerCount),
			IsFollow:      false, // TODO(vgalaxy)
		},
	}

	c.JSON(consts.StatusOK, resp)
}
