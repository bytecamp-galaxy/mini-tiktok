package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/feed"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface. get 30 latest videos with db
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// check user
	if req.UserViewId != redis.InvalidUserId {
		_, err := pack.QueryUser(ctx, req.UserViewId)
		if err != nil {
			return nil, err
		}
	}

	// if there isn't latestTime, use current time
	latestTime := req.GetLatestTime()
	if latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = curTime
	}

	// query videos in db
	v := query.Video

	// find latest 30 videos
	videos, err := v.WithContext(ctx).
		Preload(v.Author).
		Limit(conf.Init().GetInt("feed-server.default-limit")).
		Order(v.CreatedAt.Desc()).
		Where(v.CreatedAt.Lt(latestTime)).
		Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// set nextTime
	var nextTime int64
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videos[len(videos)-1].CreatedAt
	}

	// convert model.Videos to rpcmodel.Videos
	respVideos := make([]*rpcmodel.Video, len(videos))
	for i, video := range videos {
		respVideos[i], err = convert.VideoConverterORM(ctx, query.Q, video, req.UserViewId)
		if err != nil {
			return nil, err
		}
	}

	resp = &feed.FeedResponse{
		VideoList: respVideos,
		NextTime:  &nextTime,
	}
	return resp, nil
}
