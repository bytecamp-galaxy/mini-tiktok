package main

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/middlewares/client/sd"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/registry/etcd"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"time"
)

func main() {
	log.InitHLogger()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName("tiktok.client"),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	c, err := client.NewClient()
	if err != nil {
		panic(err)
	}
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	c.Use(hertztracing.ClientMiddleware(), sd.Discovery(r))

	ctx, span := otel.Tracer("github.com/hertz-contrib/obs-opentelemetry").
		Start(context.Background(), "login")
	_, b, err := c.Post(
		ctx, nil,
		"http://tiktok.api.service/douyin/user/login/?username=123456&password=ohkO4OSSw1611fR", // note no port here
		nil, config.WithSD(true))
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
		return
	}
	hlog.CtxInfof(ctx, "hertz client %s", string(b))
	span.SetAttributes(attribute.String("msg", string(b)))
	span.End()

	for {
		<-time.After(time.Second)
	}
}
