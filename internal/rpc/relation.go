package rpc

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/relation/relationservice"
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

var relationClient *relationservice.Client

func InitRelationClient(serviceName string) (*relationservice.Client, error) {
	// lazy initialization
	if relationClient != nil {
		return relationClient, nil
	}

	v := conf.Init()
	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		return nil, err
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint(fmt.Sprintf("%s:%d", v.GetString("otlp-receiver.host"), v.GetInt("otlp-receiver.port"))),
		provider.WithInsecure(),
	)
	// TODO(vgalaxy): shutdown provider

	c, err := relationservice.NewClient(
		v.GetString("relation-server.name"),
		client.WithResolver(r),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: serviceName}),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	if err != nil {
		return nil, err
	}

	relationClient = &c

	// init kitex client logger
	log.InitKLogger()

	return relationClient, nil
}
