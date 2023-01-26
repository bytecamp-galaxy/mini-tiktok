package main

import (
	user "github.com/bytecamp-galaxy/mini-tiktok/comment-server/kitex_gen/user/commentservice"
	"log"
)

func main() {
	svr := user.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
