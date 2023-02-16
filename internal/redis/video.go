package redis

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

const (
	videoIdExistKey string = "vid"
)

// LoadVideoFromDBToRedis only called by api service
func LoadVideoFromDBToRedis(ctx context.Context) error {
	// TODO: pipelined

	// init bloom filter
	// ignore recreate error
	_ = VideoIdInitBF(ctx)
	// query db
	vs, err := query.Video.WithContext(ctx).Find()
	if err != nil {
		return err
	}
	// foreach
	for _, v := range vs {
		// load video id
		err = VideoIdAddBF(ctx, v.ID)
		if err != nil {
			return err
		}
	}
	hlog.Infof("load %v video(s) from db to redis successfully", len(vs))
	return nil
}

/*==================================================================
                          Video Id
====================================================================*/

func VideoIdInitBF(ctx context.Context) error {
	return BFInit(ctx, videoIdExistKey)
}

func VideoIdExistBF(ctx context.Context, vid int64) (bool, error) {
	return BFExists(ctx, videoIdExistKey, vid)
}

func VideoIdAddBF(ctx context.Context, vid int64) error {
	return BFAdd(ctx, videoIdExistKey, vid)
}
