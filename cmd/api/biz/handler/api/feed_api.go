// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/feed"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/marmotedu/errors"
)

// GetFeed .
// @router /douyin/feed/ [GET]
func GetFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FeedRequest

	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// get the latest time.
	// if the latest time hasn't been passed as param, it's 0 by default.
	latestTime := req.GetLatestTime()
	// TODO(vgalaxy): use token

	// set up connection with feed server
	cli, err := rpc.InitFeedClient()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	// call rpc service
	reqRpc := &feed.FeedRequest{
		LatestTime: &latestTime,
	}
	respRpc, err := (*cli).GetFeed(ctx, reqRpc)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			e := errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage())
			pack.Error(c, errors.WrapC(e, errno.ErrRPCProcess, ""))
			return
		} else {
			// assume
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	// convert model.Videos to feed.Videos
	respVideos := make([]*api.Video, len(respRpc.VideoList))
	for i, video := range respRpc.VideoList {
		author := video.Author
		u := &api.User{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   &author.FollowCount,
			FollowerCount: &author.FollowerCount,
			IsFollow:      author.IsFollow,
		}
		respVideos[i] = &api.Video{
			Id:            video.Id,
			Author:        u,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
	}

	// response to client
	resp := &api.FeedResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		VideoList:  respVideos,
		NextTime:   respRpc.NextTime,
	}

	c.JSON(consts.StatusOK, resp)
}