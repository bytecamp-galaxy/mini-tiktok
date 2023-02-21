package main

import (
	"context"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/publish/handler"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/oss"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish/publishservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
	"time"
)

func main() {
	v := conf.Init()

	// init oss
	oss.Init(v.GetString("oss.type"))

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

	serverAddr := fmt.Sprintf("%s:%d", v.GetString("publish-server.host"), 0)
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("publish-server.name")),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", v.GetString("otlp-receiver.host"), v.GetInt("otlp-receiver.port"))),
		provider.WithInsecure(),
	)
	ctx := context.Background()
	defer p.Shutdown(ctx)

	var info = &registry.Info{}
	svr := publishservice.NewServer(
		new(handler.PublishServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithRegistryInfo(info),
		server.WithLimit(&limit.Option{MaxConnections: 100000, MaxQPS: 10000}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: v.GetString("publish-server.name")}),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
	)

	// fetch endpoint and init logger
	go func() {
		for {
			if info.Addr != nil {
				addr, err := net.ResolveTCPAddr("tcp", info.Addr.String())
				if err != nil {
					panic(err)
				}
				// init logger
				log.SetOutput(fmt.Sprintf("%s-%d", v.GetString("publish-server.log-path"), addr.Port))
				log.InitKLogger()
				break
			}
			time.Sleep(time.Second)
		}
	}()

	// run server
	go func() {
		err = svr.Run()
		if err != nil {
			panic(err)
		}
	}()

	// wait loop
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
		}
	}
}
