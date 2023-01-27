package main

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func main() {
	serviceName := "tiktok.client"

	log.InitHLogger()

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	c, _ := client.NewClient()
	c.Use(hertztracing.ClientMiddleware())

	for {
		ctx, span := otel.Tracer("github.com/hertz-contrib/obs-opentelemetry").
			Start(context.Background(), "loop")

		_, b, err := c.Get(ctx, nil, "http://localhost:8080/ping")
		if err != nil {
			hlog.CtxErrorf(ctx, err.Error())
		}

		span.SetAttributes(attribute.String("msg", string(b)))

		hlog.CtxInfof(ctx, "hertz client %s", string(b))
		span.End()

		<-time.After(time.Second)
	}
}
