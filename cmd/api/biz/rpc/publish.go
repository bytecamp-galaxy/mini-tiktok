package rpc

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish/publishservice"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var publishClient *publishservice.Client

func InitPublishClient() (*publishservice.Client, error) {
	// lazy initialization
	if publishClient != nil {
		return publishClient, nil
	}

	v := conf.Init()
	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		return nil, err
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("api-server.name")),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", v.GetString("otlp-receiver.host"), v.GetInt("otlp-receiver.port"))),
		provider.WithInsecure(),
	)
	// TODO(vgalaxy): shutdown provider

	c, err := publishservice.NewClient(
		v.GetString("publish-server.name"),
		client.WithResolver(r),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: v.GetString("api-server.name")}),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	if err != nil {
		return nil, err
	}

	publishClient = &c

	// init kitex client logger
	log.InitKLogger()

	return publishClient, nil
}
