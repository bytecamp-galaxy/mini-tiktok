// Code generated by hertz generator.

package main

import (
	handler "github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/pprof"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// register pprof
	pprof.Register(r)
}
