package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/feed"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface. get 30 latest videos with db
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	latestTime := req.GetLatestTime()
	uid := req.GetUserId()

	// if there isn't latestTime, use current time
	if latestTime == 0 {
		curTime := time.Now().UnixMilli()
		latestTime = curTime
	}

	// query videos in db
	q := query.Q
	v := q.Video
	u := q.User

	// find user, if not found, user is nil
	user, _ := u.WithContext(ctx).Where(u.ID.Eq(uid)).Take()

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

	var nextTime int64
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videos[len(videos)-1].CreatedAt
	}

	// convert model.Videos to rpcmodel.Videos
	respVideos := make([]*rpcmodel.Video, len(videos))
	for i, video := range videos {
		respVideos[i] = pack.VideoConverterORM(ctx, q, video, user)
	}

	resp = &feed.FeedResponse{
		VideoList: respVideos,
		NextTime:  &nextTime,
	}
	return resp, nil
}
