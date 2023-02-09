package main

import (
	"context"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/publish/handler"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish/publishservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/minio"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	v := conf.Init()
	// init log
	log.SetOutput(v.GetString("publish-server.log-path"))
	log.InitKLogger()

	// init minio
	minio.Init()

	// init db
	dal.Init(false)

	// init redis
	redis.Init()

	// init snowflake id generator
	snowflake.Init()

	// init server

	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	serverAddr := fmt.Sprintf("%s:%d", v.GetString("publish-server.host"), v.GetInt("publish-server.port"))
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("publish-server.name")),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", v.GetString("otlp-receiver.host"), v.GetInt("otlp-receiver.port"))),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: v.GetString("publish-server.name")}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	// run server
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
