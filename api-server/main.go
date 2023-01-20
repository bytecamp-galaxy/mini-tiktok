// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network/netpoll"
	"mini-tiktok-v2/api-server/biz/middleware"
	"mini-tiktok-v2/api-server/biz/registry/eureka"
	"mini-tiktok-v2/pkg/dal"
	"time"
)

func Init() {
	dal.Init()
	middleware.Init()
}

func main() {
	Init()
	addr := "localhost:8080"
	r := eureka.NewEurekaRegistry([]string{"http://localhost:8761/eureka"}, 40*time.Second)
	h := server.Default(server.WithHostPorts(addr),
		server.WithTransport(netpoll.NewTransporter),
		server.WithExitWaitTime(3*time.Second),
		server.WithRegistry(r, &registry.Info{
			ServiceName: "tiktok.api.service",
			Addr:        utils.NewNetAddr("tcp", addr),
			Weight:      10,
			Tags:        nil,
		}))
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		fmt.Println("before ctx.Done()")
		<-ctx.Done()
		fmt.Println("after ctx.Done()")
	})
	register(h)
	h.Spin()
}
