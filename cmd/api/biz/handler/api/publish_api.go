// Code generated by hertz generator.

package api

import (
	"bytes"
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/jwt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/marmotedu/errors"
	"io"
)

// PublishAction .
// @router /douyin/publish/action/ [POST]
func PublishAction(ctx context.Context, c *app.RequestContext) {
	// NOTE: DO NOT USE `BindAndValidate`
	fileHeader, err := c.Request.FormFile("data")
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrBindAndValidation, err.Error()))
		return
	}

	// get .mp4 data
	file, err := fileHeader.Open()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, err.Error()))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, err.Error()))
		return
	}

	// fetch user id from token
	userId, ok := c.Get(jwt.IdentityKey)
	if !ok {
		pack.Error(c, errors.WithCode(errno.ErrUnknown, pack.BrokenInvariantStatusMessage))
		return
	}

	// set up connection with publish server
	cli, err := rpc.InitPublishClient()
	if err != nil {
		pack.Error(c, errors.WithCode(errno.ErrClientRPCInit, err.Error()))
		return
	}

	// call rpc service
	reqRpc := &publish.PublishRequest{
		UserId: userId.(int64),
		Data:   buf.Bytes(),
		Title:  c.PostForm("title"),
	}

	_, err = (*cli).PublishVideo(ctx, reqRpc)
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

	resp := &api.PublishActionResponse{
		StatusCode: errno.ErrSuccess,
		StatusMsg:  utils.String(pack.SuccessStatusMessage),
	}

	c.JSON(consts.StatusOK, resp)
}
