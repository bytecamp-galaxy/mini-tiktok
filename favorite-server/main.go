package main

import (
	favorite "github.com/bytecamp-galaxy/mini-tiktok/favorite-server/kitex_gen/favorite/favoriteservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/registry-eureka/registry"
	"net"
	"time"
)

func main() {
	dal.Init()
	r := registry.NewEurekaRegistry([]string{"http://81.68.219.146:8761/eureka"}, 3*time.Second)
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	svr := favorite.NewServer(new(FavoriteServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "tiktok.favorite.service",
		}),
		server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
