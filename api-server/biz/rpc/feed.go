package rpc

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/feed-server/kitex_gen/feed/feedservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient *feedservice.Client

func InitFeedClient() (*feedservice.Client, error) {
	// lazy initialization
	if feedClient != nil {
		return feedClient, nil
	}

	v := conf.Init().V
	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		return nil, err
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("feed-server.name")),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	// TODO(vgalaxy): shutdown provider

	c, err := feedservice.NewClient(
		v.GetString("feed-server.name"),
		client.WithResolver(r),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: v.GetString("api-server.name")}),
	)
	if err != nil {
		return nil, err
	}

	feedClient = &c

	// init kitex client logger
	log.InitKLogger()

	return feedClient, nil
}
