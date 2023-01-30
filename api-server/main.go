// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/jwt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/registry/etcd"
	"net"
	"time"
)

func main() {
	// init db
	dal.Init()

	// init jwt
	jwt.Init()

	// init errno
	errno.Init()

	// init log
	log.InitHLogger()

	// init server
	v := conf.Init().V

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
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := tracing.NewServerTracer()
	h := server.Default(
		server.WithHostPorts(serverAddr),
		server.WithTransport(netpoll.NewTransporter),
		server.WithExitWaitTime(time.Duration(v.GetInt("api-server.exit-wait-time"))*time.Second),
		server.WithRegistry(r, &registry.Info{
			ServiceName: v.GetString("api-server.name"),
			Addr:        addr,
		}),
		tracer)
	h.Use(tracing.ServerMiddleware(cfg))

	// register
	register(h)

	// run server
	h.Spin()
}
