package main

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
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

	c, _ := client.NewClient()
	c.Use(hertztracing.ClientMiddleware())

	ctx, span := otel.Tracer("github.com/hertz-contrib/obs-opentelemetry").
		Start(context.Background(), "register")
	_, b, err := c.Post(ctx, nil, "http://localhost:8080/douyin/user/register/?username=123456&password=ohkO4OSSw1611fR", nil)
	if err != nil {
		hlog.CtxErrorf(ctx, err.Error())
	}
	hlog.CtxInfof(ctx, "hertz client %s", string(b))
	span.SetAttributes(attribute.String("msg", string(b)))
	span.End()

	for {
		<-time.After(time.Second)
	}
}
