package main

import (
	"github.com/cloudwego/kitex/server"
	"mini-tiktok-v2/pkg/dal"
	user "mini-tiktok-v2/user-server/kitex_gen/user/userservice"
	"net"
)

func main() {
	dal.Init()
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:8888")
	if err != nil {
		panic(err)
	}
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr))
	err = svr.Run()
	if err != nil {
		panic(err)
	}
}
