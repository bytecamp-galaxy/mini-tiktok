// Code generated by hertz generator.

package api

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/comment"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/marmotedu/errors"
)

// CommentAction .
// @router /douyin/comment/action/ [POST]
// @description 评论操作：登录用户对视频进行评论
// @produce application/json
// @param q query api.CommentActionRequest true "comment action request"
// @success 200 {object} api.CommentActionResponse
func CommentAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentActionRequest

	// bind and validate request.
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// fetch user id from token
	userId, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrParseToken, ""))
		return
	}

	reqRPC := &comment.CommentActionRequest{
		VideoId:     req.VideoId,
		UserId:      userId.(int64),
		ActionType:  req.ActionType,
		CommentText: req.CommentText,
		CommentId:   req.CommentId,
	}

	// set up connection with comment server
	v := conf.Init()
	cli, err := rpc.InitCommentClient(v.GetString("api-server.name"))
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	respRPC, err := (*cli).CommentAction(ctx, reqRPC)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			pack.Error(c, errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage()))
			return
		} else {
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	resp := &api.CommentActionResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
		Comment:    convert.CommentConverterAPI(respRPC.Comment), // maybe nil
	}

	c.JSON(consts.StatusOK, resp)
}

// CommentList .
// @router /douyin/comment/list [GET]
// @description 评论列表：查看视频的所有评论，按发布时间倒序
// @produce application/json
// @param q query api.CommentListRequest true "comment list request"
// @success 200 {object} api.CommentListResponse
func CommentList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req api.CommentListRequest

	// bind and validate request.
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// fetch user view id from token
	userViewId, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrParseToken, ""))
		return
	}

	// set up connection with comment server
	v := conf.Init()
	cli, err := rpc.InitCommentClient(v.GetString("api-server.name"))
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	reqRPC := comment.CommentListRequest{
		VideoId:    req.VideoId,
		UserViewId: userViewId.(int64),
	}
	respRPC, err := (*cli).CommentList(ctx, &reqRPC)
	if err != nil {
		if bizErr, ok := kerrors.FromBizStatusError(err); ok {
			pack.Error(c, errors.WithCode(int(bizErr.BizStatusCode()), bizErr.BizMessage()))
			return
		} else {
			pack.Error(c, errors.WithCode(errno.ErrRPCLink, err.Error()))
			return
		}
	}

	list := make([]*api.Comment, len(respRPC.CommentList))
	for i, c := range respRPC.CommentList {
		list[i] = convert.CommentConverterAPI(c)
	}

	resp := &api.CommentListResponse{
		StatusCode:  errno.ErrSuccess,
		StatusMsg:   utils.String(pack.SuccessStatusMessage),
		CommentList: list,
	}

	c.JSON(consts.StatusOK, resp)
}
