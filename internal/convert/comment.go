package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"time"
)

// CommentConverterAPI convert *rpcmodel.Comment to *api.Comment, can only be called by api servers
func CommentConverterAPI(comment *rpcmodel.Comment) *api.Comment {
	if comment == nil {
		return nil
	}
	return &api.Comment{
		Id:         comment.Id,
		User:       UserConverterAPI(comment.User),
		Content:    comment.Content,
		CreateDate: comment.CreateDate,
	}
}

// CommentConverterORM convert *model.Comment to *rpcmodel.Comment, can only be called by rpc servers
func CommentConverterORM(ctx context.Context, q *query.Query, comment *model.Comment, userViewId int64) (res *rpcmodel.Comment, err error) {
	if comment == nil {
		return nil, nil
	}
	user, err := UserConverterORM(ctx, q, &comment.User, userViewId) // preload required
	if err != nil {
		return nil, err
	}
	return &rpcmodel.Comment{
		Id:         comment.ID,
		User:       user,
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreatedAt, 0).Format("01-02"),
	}, nil
}
