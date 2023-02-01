package main

import (
	"context"
	"fmt"
	comment "github.com/bytecamp-galaxy/mini-tiktok/comment-server/kitex_gen/comment/commentservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"net"
)

func main() {
	// init db
	dal.Init()

	// init log
	log.InitKLogger()

	// init server
	v := conf.Init()

	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdRegistry([]string{etcdAddr})
	if err != nil {
		panic(err)
	}

	serverAddr := fmt.Sprintf("%s:%d", v.GetString("comment-server.host"), v.GetInt("comment-server.port"))
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("comment-server.name")),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	svr := comment.NewServer(
		new(CommentServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithMuxTransport(),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: v.GetString("comment-server.name")}),
	)

	// run server
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
