package main

import (
	user "github.com/bytecamp-galaxy/mini-tiktok/comment-server/kitex_gen/user/commentservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-eureka/registry"
	"net"
	"time"
)

func main() {
	dal.Init()
	r := registry.NewEurekaRegistry([]string{"http://localhost:8761/eureka"}, 3*time.Second)
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(CommentServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "tiktok.comment.service",
		}),
		server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
