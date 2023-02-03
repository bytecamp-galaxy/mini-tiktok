// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/marmotedu/errors"
)

// FavoriteAction .
// @router /douyin/favorite/action/ [POST]
func FavoriteAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	uid, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, pack.BrokenInvariantStatusMessage))
		return
	}

	reqRPC := &favorite.FavoriteActionRequest{
		UserId:     uid.(int64),
		VideoId:    req.VideoId,
		ActionType: req.ActionType,
	}

	v := conf.Init()
	cli, err := rpc.InitFavoriteClient(v.GetString("api-server.name"))
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	_, err = (*cli).FavoriteAction(ctx, reqRPC)
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

	resp := &api.FavoriteActionResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
	}

	c.JSON(consts.StatusOK, resp)
}

// FavoriteList .
// @router /douyin/favorite/list/ [GET]
func FavoriteList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.FavoriteListRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	_, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, pack.BrokenInvariantStatusMessage))
		return
	}

	// check user id
	//if uid != req.UserId {
	//	fmt.Println(uid, req.UserId)
	//	pack.Error(c, errors.WithCode(errno.ErrTokenInvalid, "inconsistent user id"))
	//	return
	//}

	// set up connection with comment server
	v := conf.Init()
	cli, err := rpc.InitFavoriteClient(v.GetString("api-server.name"))
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	reqRPC := favorite.FavoriteListRequest{
		UserId: req.UserId,
	}

	respRPC, err := (*cli).FavoriteList(ctx, &reqRPC)
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

	list := make([]*api.Video, len(respRPC.VideoList))

	for i, video := range respRPC.VideoList {
		author := video.Author
		u := &api.User{
			Id:            author.Id,
			Name:          author.Name,
			FollowCount:   &author.FollowCount,
			FollowerCount: &author.FollowerCount,
			IsFollow:      author.IsFollow,
		}
		list[i] = &api.Video{
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

	resp := &api.FavoriteListResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		VideoList:  list,
	}

	c.JSON(consts.StatusOK, resp)
}
