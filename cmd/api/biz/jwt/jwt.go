package jwt

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
	"time"
)

var (
	Middleware  *jwt.HertzJWTMiddleware
	IdentityKey = "id"
)

func Init() {
	var err error
	Middleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Realm:   "mini-tiktok",
		Key:     []byte("cyvG2OzO9KQNsY3"),
		Timeout: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok { // save user id
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		TokenLookup: "query:token",
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": errno.ErrUnauthorized,
				"status_msg":  message,
			})
		},
		IdentityKey: IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c) // load user id
			if v, ok := claims[IdentityKey]; ok {
				return int64(v.(float64))
			}
			return 0 // return invalid id
		},
	})
	if err != nil {
		panic("JWT Error:" + err.Error())
	}
}
