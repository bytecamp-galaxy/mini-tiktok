package main

import (
	"context"
	feed "github.com/bytecamp-galaxy/mini-tiktok/feed-server/kitex_gen/feed"
	"github.com/bytecamp-galaxy/mini-tiktok/feed-server/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"time"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

const (
	LIMIT = 30 // 单次返回最大视频数
)

// GetFeed implements the FeedServiceImpl interface. get 30 latest videos with db
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {

	latestTime := req.GetLatestTime()
	// if there isn't  latestTime, use current time
	if latestTime == 0 {
		curTime := int64(time.Now().UnixMilli())
		latestTime = curTime
	}

	// query videos in db
	q := query.Q
	v := q.Video

	// find latest 30 videos
	videos, err := v.WithContext(ctx).Limit(LIMIT).Order(v.CreatedAt.Desc()).Where(v.CreatedAt.Lt(latestTime)).Find()

	if err != nil {
		return nil, err
	}

	var nextTime int64
	if len(videos) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videos[len(videos)-1].CreatedAt
	}

	// convert model.Videos to feed.Videos
	respVideos := make([]*feed.Video, len(videos))
	for i, video := range videos {
		author := video.Author
		u := &user.User{
			Id:            author.ID,
			Name:          author.Username,
			FollowCount:   author.FollowingCount,
			FollowerCount: author.FollowerCount,
			IsFollow:      false, // TODO
		}
		respVideos[i] = &feed.Video{
			Id:            video.ID,
			Author:        u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false, // TODO: 如果用户登录状态下刷视频，如何高效的获取这些用户对刷到的视频的点赞信息？
			Title:         video.Title,
		}
	}

	statusCode := "success"
	resp = &feed.FeedResponse{
		StatusCode: 0,
		StatusMsg:  &statusCode,
		VideoList:  respVideos,
		NextTime:   &nextTime,
	}
	return resp, nil
}
