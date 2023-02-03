package handler

import (
	"context"
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
	user, _ := u.WithContext(ctx).Where(u.ID.Eq(req.GetUserId())).Take()

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
		isFavorite := false
		// TODO: 如果用户登录状态下刷视频，如何高效的获取这些用户对刷到的视频的点赞信息？
		if user != nil && u.FavoriteVideos.WithContext(ctx).Where(v.ID.Eq(video.ID)).Model(user).Count() != 0 {
			isFavorite = true
		}
		author := video.Author
		u := &rpcmodel.User{
			Id:            author.ID,
			Name:          author.Username,
			FollowCount:   author.FollowingCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      false, // TODO
		}
		respVideos[i] = &rpcmodel.Video{
			Id:            video.ID,
			Author:        u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		}
	}

	resp = &feed.FeedResponse{
		VideoList: respVideos,
		NextTime:  &nextTime,
	}
	return resp, nil
}
