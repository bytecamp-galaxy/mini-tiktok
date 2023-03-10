// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	_ "github.com/bytecamp-galaxy/mini-tiktok/cmd/api/docs"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/registry/etcd"
	"net"
	"time"
)

// @title mini tiktok
// @version 1.0
// @description 第五届字节跳动青训营后端专场结营项目 - 极简版抖音

// @contact.name bytecamp-galaxy
// @contact.url https://github.com/bytecamp-galaxy/

// @license.name MIT License
// @license.url https://mit-license.org/

// @host localhost:8080
// @BasePath /douyin/
// @schemes http
func main() {
	v := conf.Init()
	// init log
	log.SetOutput(v.GetString("api-server.log-path"))
	log.InitHLogger()

	// init db
	dal.Init(true)

	// init redis
	redis.Init()
	err := redis.LoadUserFromDBToRedis(context.Background())
	if err != nil {
		panic(err)
	}
	err = redis.LoadVideoFromDBToRedis(context.Background())
	if err != nil {
		panic(err)
	}

	// init jwt
	jwt.Init()

	// init errno
	// NOTE: only register error code when api server setup, since only parse error in api server
	errno.Init()

	// init server
	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	serverAddr := fmt.Sprintf("%s:%d", v.GetString("api-server.host"), v.GetInt("api-server.port"))
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("api-server.name")),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", v.GetString("otlp-receiver.host"), v.GetInt("otlp-receiver.port"))),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := tracing.NewServerTracer()
	h := server.New(
		server.WithHostPorts(serverAddr),
		server.WithStreamBody(true),
		server.WithTransport(standard.NewTransporter),
		server.WithExitWaitTime(time.Duration(v.GetInt("api-server.exit-wait-time"))*time.Second),
		server.WithMaxRequestBodySize(v.GetInt("api-server.max-request-body-size")),
		server.WithRegistry(r, &registry.Info{
			ServiceName: v.GetString("api-server.name"),
			Addr:        addr,
		}),
		tracer)

	// set global middleware
	h.Use(
		// tracer
		tracing.ServerMiddleware(cfg),
		// recovery
		recovery.Recovery(recovery.WithRecoveryHandler(
			func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
				hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
				c.JSON(consts.StatusInternalServerError, utils.H{
					"status_code": errno.ErrStatusInternalServerError,
					"status_msg":  fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
				})
			},
		)),
		// gzip
		gzip.Gzip(gzip.DefaultCompression),
		// access log
		func(c context.Context, ctx *app.RequestContext) {
			start := time.Now()
			ctx.Next(c)
			end := time.Now()
			hlog.Infof("status=%d cost=%s method=%s full_path=%s client_ip=%s host=%s",
				ctx.Response.StatusCode(), end.Sub(start).String(),
				ctx.Request.Header.Method(), ctx.Request.URI().String(), ctx.ClientIP(), ctx.Request.Host())
		},
	)

	// set NoRoute handler
	h.NoRoute(func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusNotFound, map[string]interface{}{
			"status_code": errno.ErrStatusNotFound,
			"status_msg":  "no route",
		})
	})

	// set NoMethod handler
	h.NoMethod(func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusMethodNotAllowed, map[string]interface{}{
			"status_code": errno.ErrStatusMethodNotAllowed,
			"status_msg":  "no method",
		})
	})

	// register
	register(h)

	// run server
	h.Spin()
}
