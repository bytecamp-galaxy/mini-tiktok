// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/mw"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/bytecamp-galaxy/mini-tiktok/user-server/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/user-server/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-eureka/resolver"

	"github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/model/api"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest

	// bind and validate request
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &api.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// validate password
	err = utils.ValidatePassword(req.Password)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &api.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// set up connection with user server
	r := resolver.NewEurekaResolver([]string{"http://localhost:8761/eureka"})
	cli, err := userservice.NewClient("tiktok.user.service", client.WithResolver(r))
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// call rpc service
	reqRpc := &user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := cli.UserRegister(ctx, reqRpc)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// handle status code
	if respRpc.StatusCode != 0 {
		c.JSON(consts.StatusInternalServerError, &api.UserRegisterResponse{
			StatusCode: respRpc.StatusCode,
			StatusMsg:  utils.String(respRpc.StatusMsg),
		})
		return
	}

	// generate token
	token, _, err := mw.JwtMiddleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserRegisterResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// response to client
	resp := &api.UserRegisterResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  utils.String(respRpc.StatusMsg),
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

	// bind and validate request
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &api.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// set up connection with user server
	r := resolver.NewEurekaResolver([]string{"http://localhost:8761/eureka"})
	cli, err := userservice.NewClient("tiktok.user.service", client.WithResolver(r))
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// call rpc service
	reqRpc := &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := cli.UserLogin(ctx, reqRpc)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// handle status code
	if respRpc.StatusCode != 0 {
		c.JSON(consts.StatusInternalServerError, &api.UserLoginResponse{
			StatusCode: respRpc.StatusCode,
			StatusMsg:  utils.String(respRpc.StatusMsg),
		})
		return
	}

	// generate token
	token, _, err := mw.JwtMiddleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserLoginResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// response to client
	resp := &api.UserLoginResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  utils.String(respRpc.StatusMsg),
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

	// bind and validate request
	err = c.BindAndValidate(&req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, &api.UserQueryResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// fetch user id from token
	id, ok := c.Get(mw.IdentityKey)
	if !ok {
		c.JSON(consts.StatusInternalServerError, &api.UserQueryResponse{
			StatusCode: 1,
			StatusMsg:  utils.String("broken invariant"),
		})
		return
	}

	// check user id
	if id != req.UserId {
		c.JSON(consts.StatusUnauthorized, &api.UserQueryResponse{
			StatusCode: 1,
			StatusMsg:  utils.String("incorrect id"),
		})
		return
	}

	// set up connection with user server
	r := resolver.NewEurekaResolver([]string{"http://localhost:8761/eureka"})
	cli, err := userservice.NewClient("tiktok.user.service", client.WithResolver(r))
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserQueryResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// call rpc service
	reqRpc := &user.UserQueryRequest{
		UserId: req.UserId,
	}

	respRpc, err := cli.UserQuery(ctx, reqRpc)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, &api.UserQueryResponse{
			StatusCode: 1,
			StatusMsg:  utils.String(err.Error()),
		})
		return
	}

	// handle status code
	if respRpc.StatusCode != 0 {
		c.JSON(consts.StatusInternalServerError, &api.UserQueryResponse{
			StatusCode: respRpc.StatusCode,
			StatusMsg:  utils.String(respRpc.StatusMsg),
		})
		return
	}

	// response to client
	resp := &api.UserQueryResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  utils.String(respRpc.StatusMsg),
		User: &api.User{
			Id:            respRpc.User.Id,
			Name:          respRpc.User.Name,
			FollowCount:   utils.Int64(respRpc.User.FollowerCount),
			FollowerCount: utils.Int64(respRpc.User.FollowerCount),
			IsFollow:      false, // TODO
		},
	}

	c.JSON(consts.StatusOK, resp)
}
