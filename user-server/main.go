package main

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	user "github.com/bytecamp-galaxy/mini-tiktok/user-server/kitex_gen/user/userservice"
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
	v := conf.Init().V

	eurekaAddr := fmt.Sprintf("http://%s:%d/eureka", v.GetString("eureka.host"), v.GetInt("eureka.port"))
	interval := v.GetInt64("eureka.rpc-heartbeat-interval")
	r := registry.NewEurekaRegistry([]string{eurekaAddr}, time.Duration(interval)*time.Second)

	serverAddr := fmt.Sprintf("%s:%d", v.GetString("user-server.host"), v.GetInt("user-server.port"))
	addr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		panic(err)
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: v.GetString("user-server.name"),
		}),
		server.WithServiceAddr(addr))

	// init log
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	// run server
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
