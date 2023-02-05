package convert

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/model/api"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"time"
)

// CommentConverterAPI convert *rpcmodel.Comment to *api.Comment
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

// CommentConverterORM convert *model.Comment to *rpcmodel.Comment
func CommentConverterORM(ctx context.Context, q *query.Query, comment *model.Comment, view *model.User) *rpcmodel.Comment {
	if comment == nil {
		return nil
	}
	user := UserConverterORM(ctx, q, &comment.User, view) // preload required
	return &rpcmodel.Comment{
		Id:         comment.ID,
		User:       user,
		Content:    comment.Content,
		CreateDate: time.Unix(comment.CreatedAt, 0).String(),
	}
}
