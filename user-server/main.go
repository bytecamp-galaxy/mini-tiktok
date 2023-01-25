package main

import (
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
	r := registry.NewEurekaRegistry([]string{"http://localhost:8761/eureka"}, 3*time.Second)
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl),
		server.WithMiddleware(mw.CommonMiddleware),
		server.WithMiddleware(mw.ServerMiddleware),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "tiktok.user.service",
		}),
		server.WithServiceAddr(addr))

	// init log
	//f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer func(f *os.File) {
	//	err := f.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(f)
	//hlog.SetOutput(f)
	klog.SetLogger(kitexzap.NewLogger())
	klog.SetLevel(klog.LevelDebug)

	// run server
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
