package main

import (
	feed "github.com/bytecamp-galaxy/mini-tiktok/feed-server/kitex_gen/feed/feedservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/constants"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"github.com/kitex-contrib/registry-eureka/registry"
	"net"
	"time"
)

func main() {
	dal.Init()

	// init server
	r := registry.NewEurekaRegistry([]string{constants.EurekaServerUrl}, 3*time.Second)
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	svr := feed.NewServer(new(FeedServiceImpl),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "tiktok.feed.service",
		}),
		server.WithServiceAddr(addr))

	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	// run server
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
