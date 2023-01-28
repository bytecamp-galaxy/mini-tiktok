package main

import (
	"context"
	user "github.com/bytecamp-galaxy/mini-tiktok/comment-server/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *user.CommentActionRequest) (resp *user.CommentActionResponse, err error) {
	q := query.Use(mysql.DB)
	err = q.Transaction(func(tx *query.Query) error {
		switch req.ActionType {
		case 1:
			{
				c := model.Comment{
					VideoID: req.VideoId,
					UserID:  req.UserId,
					Content: *req.CommentText,
				}
				err = tx.Comment.WithContext(ctx).Create(&c)
				if err != nil {
					return err
				}
				v := tx.Video
				_, err = v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Add(1))
				if err != nil {
					return err
				}

				resp = &user.CommentActionResponse{
					StatusCode: 0,
					Comment: &user.Comment{
						Id:         c.ID,
						User:       nil,
						Conent:     c.Content,
						CreateDate: time.Unix(c.CreatedAt, 0).String(),
					},
				}
			}
		case 2:
			{
				c := tx.Comment
				_, err = c.WithContext(ctx).Where(c.ID.Eq(*req.CommentId)).Delete()
				if err != nil {
					return err
				}

				v := tx.Video
				_, err = v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Sub(1))
				if err != nil {
					return err
				}

				resp = &user.CommentActionResponse{
					StatusCode: 0,
				}
			}
		default:
			panic("Request argument violates convention")
		}
		return nil
	})

	if err != nil {
		resp = &user.CommentActionResponse{
			StatusCode: -1,
		}
		return resp, err
	}
	return resp, err
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *user.CommentListRequest) (resp *user.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}
