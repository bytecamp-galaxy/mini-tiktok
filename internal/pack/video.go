package pack

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// QueryVideo can only be called by rpc servers
func QueryVideo(ctx context.Context, vid int64) (*model.Video, error) {
	// query video id in redis bloom filter
	exist, err := redis.VideoIdExistBF(ctx, vid)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}
	if !exist {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidVideo), "")
	}

	// query video in db
	v, _ := query.Video.WithContext(ctx).Where(query.Video.ID.Eq(vid)).Take()
	if v == nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidVideo), "")
	}

	return v, nil
}
