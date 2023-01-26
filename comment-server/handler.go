package main

import (
	"context"
	user "github.com/bytecamp-galaxy/mini-tiktok/comment-server/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *user.CommentActionRequest) (resp *user.CommentActionResponse, err error) {
	q := query.Use(mysql.DB)

	err = q.Transaction(func(tx *query.Query) error {
		if _, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(100)).Delete(); err != nil {
			return err
		}
		if _, err := tx.Article.WithContext(ctx).Create(&model.User{Name: "modi"}); err != nil {
			return err
		}
		return nil
	})
	return
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *user.CommentListRequest) (resp *user.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
