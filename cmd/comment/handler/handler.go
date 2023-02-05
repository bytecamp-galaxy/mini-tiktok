package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/comment"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
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
				id := snowflake.Generate()
				err = tx.Comment.WithContext(ctx).Create(&model.Comment{
					ID:      id,
					VideoID: req.VideoId,
					UserID:  req.UserId,
					Content: *req.CommentText,
				})
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				v := tx.Video
				_, err = v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Add(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				c, err := tx.Comment.Preload(tx.Comment.User).WithContext(ctx).Where(tx.Comment.ID.Eq(id)).Take()
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				resp = &comment.CommentActionResponse{
					Comment: convert.CommentConverterORM(ctx, q, c, nil), // 不允许自己关注自己
				}
			}
		case 2:
			{
				c := tx.Comment
				result, err := c.WithContext(ctx).Where(c.ID.Eq(*req.CommentId)).Delete()
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected == 0 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "nonexistent comment")
				}

				v := tx.Video
				_, err = v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Sub(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				resp = &comment.CommentActionResponse{}
			}
		default:
			return kerrors.NewBizStatusError(int32(errno.ErrUnknown), "unknown action type")
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// CommentList implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentList(ctx context.Context, req *comment.CommentListRequest) (resp *comment.CommentListResponse, err error) {
	q := query.Use(mysql.DB)
	err = q.Transaction(func(tx *query.Query) error {
		comments, err := tx.Comment.WithContext(ctx).
			Preload(tx.Comment.User).
			Where(tx.Comment.VideoID.Eq(req.VideoId)).
			Find()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		view, err := tx.User.WithContext(ctx).Where(tx.User.ID.Eq(req.UserViewId)).Take()
		if err != nil {
			return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
		}

		list := make([]*rpcmodel.Comment, len(comments))
		for i, c := range comments {
			list[i] = convert.CommentConverterORM(ctx, q, c, view)
		}
		resp = &comment.CommentListResponse{
			CommentList: list,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return resp, nil
}
