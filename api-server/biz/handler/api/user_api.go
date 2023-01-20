// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-eureka/resolver"
	"google.golang.org/protobuf/proto"
	"mini-tiktok-v2/api-server/biz/middleware"
	"mini-tiktok-v2/user-server/kitex_gen/user"
	"mini-tiktok-v2/user-server/kitex_gen/user/userservice"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	api "mini-tiktok-v2/api-server/biz/model/api"
)

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		panic(nil)
	}

	r := resolver.NewEurekaResolver([]string{"http://localhost:8761/eureka"})
	cli := userservice.MustNewClient("tiktok.user.service", client.WithResolver(r))

	reqRpc := &user.UserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := cli.UserRegister(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	token, _, err := middleware.JwtMiddleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		panic(err)
	}

	resp := &api.UserRegisterResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  proto.String(respRpc.StatusMsg),
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
		panic(nil)
	}

	r := resolver.NewEurekaResolver([]string{"http://localhost:8761/eureka"})
	cli := userservice.MustNewClient("tiktok.user.service", client.WithResolver(r))

	reqRpc := &user.UserLoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	respRpc, err := cli.UserLogin(ctx, reqRpc)
	if err != nil {
		panic(err)
	}

	token, _, err := middleware.JwtMiddleware.TokenGenerator(respRpc.UserId)
	if err != nil {
		panic(err)
	}

	resp := &api.UserLoginResponse{
		StatusCode: respRpc.StatusCode,
		StatusMsg:  proto.String(respRpc.StatusMsg),
		UserId:     respRpc.UserId,
		Token:      token,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserQuery .
// @router /douyin/user/ [GET]
func UserQuery(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.UserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		panic(err)
	}

	id, ok := c.Get(middleware.IdentityKey)
	if !ok {
		panic(err)
	}

	if id != req.UserId {
		panic(err)
	}

	resp := new(api.UserResponse)

	c.JSON(consts.StatusOK, resp)
}
