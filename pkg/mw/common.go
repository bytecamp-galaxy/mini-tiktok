package mw

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware common mw print some rpc info, real request and real response
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)
		// get real request
		// TODO(vgalaxy): record large request and response, set max length now
		klog.Infof("real request: %+.512v", req)
		// get remote service information
		klog.Infof("remote service name: %s, remote method: %s", ri.To().ServiceName(), ri.To().Method())
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		// get real response
		klog.Infof("real response: %+.512v", resp)
		return nil
	}
}
