package rpc

import (
	"fmt"
	etcd "github.com/bytecamp-galaxy/kitex-registry-etcd"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/mw"
	"github.com/bytecamp-galaxy/mini-tiktok/user-server/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

var userClient *userservice.Client

func InitUserClient() (*userservice.Client, error) {
	// lazy initialization
	if userClient != nil {
		return userClient, nil
	}

	v := conf.Init().V
	etcdAddr := fmt.Sprintf("%s:%d", v.GetString("etcd.host"), v.GetInt("etcd.port"))
	r, err := etcd.NewEtcdResolver([]string{etcdAddr})
	if err != nil {
		return nil, err
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(v.GetString("api-server.name")),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	// TODO(vgalaxy): shutdown provider

	c, err := userservice.NewClient(
		v.GetString("user-server.name"),
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

	userClient = &c

	// init kitex client logger
	log.InitKLogger()

	return userClient, nil
}
