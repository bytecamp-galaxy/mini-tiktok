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
	q.Transaction(func(tx *query.Query) error {
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
						Content:    c.Content,
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
			StatusCode: -1, //TODO(heiyan): return more meaningful status code.
		}
		return resp, err
	}
	return resp, err
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *user.CommentListRequest) (resp *user.CommentListResponse, err error) {
	q := query.Use(mysql.DB)
	q.Transaction(func(tx *query.Query) error {
		c := tx.Comment
		var comments []*model.Comment
		comments, err = c.WithContext(ctx).Where(c.VideoID.Eq(req.VideoId)).Find()

		if err != nil {
			return err
		}

		list := make([]*user.Comment, len(comments))
		for i, commentPO := range comments {
			commentVO := Po2voComment(*commentPO)
			list[i] = &commentVO
		}
		resp = &user.CommentListResponse{
			StatusCode:  0,
			StatusMsg:   nil,
			CommentList: list,
		}
		return err
	})

	if err != nil {
		return &user.CommentListResponse{
			StatusCode:  -1,
			StatusMsg:   nil,
			CommentList: nil,
		}, err
	}

	return resp, err
}

func Po2voComment(commentPO model.Comment) user.Comment {
	userVO := Po2voUser(commentPO.User)
	return user.Comment{
		Id:         commentPO.ID,
		User:       &userVO,
		Content:    commentPO.Content,
		CreateDate: time.Unix(commentPO.CreatedAt, 0).String(),
	}
}

func Po2voUser(userPO model.User) user.User {
	return user.User{
		Id:            userPO.ID,
		Name:          userPO.Username,
		FollowCount:   userPO.FollowingCount,
		FollowerCount: userPO.FollowerCount,
		IsFollow:      false,
	}
}
