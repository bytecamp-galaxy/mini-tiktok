// Code generated by Kitex v0.4.4. DO NOT EDIT.
package feedservice

import (
	feed "github.com/bytecamp-galaxy/mini-tiktok/scripts/kitex_gen/feed"
	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler feed.FeedService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
