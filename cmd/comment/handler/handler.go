package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/comment"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"time"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
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

				u := tx.User
				data, err := u.WithContext(ctx).
					Where(u.ID.Eq(req.UserId)).
					Take()
				if err != nil {
					return err
				}

				resp = &comment.CommentActionResponse{
					Comment: &rpcmodel.Comment{
						Id:         c.ID,
						User:       Po2voUser(*data),
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

				resp = &comment.CommentActionResponse{}
			}
		default:
			return kerrors.NewBizStatusError(int32(errno.ErrUnknown), "request argument violates convention")
		}
		return nil
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	q := query.Use(mysql.DB)
	err = q.Transaction(func(tx *query.Query) error {
		c := tx.Comment
		var comments []*model.Comment
		comments, err = c.WithContext(ctx).
			Preload(c.User).
			Where(c.VideoID.Eq(req.VideoId)).
			Find()

		if err != nil {
			return err
		}

		list := make([]*rpcmodel.Comment, len(comments))
		for i, commentPO := range comments {
			commentVO := Po2voComment(*commentPO)
			list[i] = &commentVO
		}
		resp = &comment.CommentListResponse{
			CommentList: list,
		}
		return err
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}
	return resp, nil
}

func Po2voComment(commentPO model.Comment) rpcmodel.Comment {
	userVO := Po2voUser(commentPO.User)
	return rpcmodel.Comment{
		Id:         commentPO.ID,
		User:       userVO,
		Content:    commentPO.Content,
		CreateDate: time.Unix(commentPO.CreatedAt, 0).String(),
	}
}

func Po2voUser(userPO model.User) *rpcmodel.User {
	return &rpcmodel.User{
		Id:            userPO.ID,
		Name:          userPO.Username,
		FollowCount:   userPO.FollowingCount,
		FollowerCount: userPO.FollowerCount,
		IsFollow:      false,
	}
}
