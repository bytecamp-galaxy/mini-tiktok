package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/comment"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// CommentServiceImpl implements the last service interface defined in the IDL.
type CommentServiceImpl struct{}

// CommentAction implements the CommentServiceImpl interface.
func (s *CommentServiceImpl) CommentAction(ctx context.Context, req *comment.CommentActionRequest) (resp *comment.CommentActionResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// check video
	_, err = pack.QueryVideo(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	// do action
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
				result, err := v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Add(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
				}

				c, err := tx.Comment.Preload(tx.Comment.User).WithContext(ctx).Where(tx.Comment.ID.Eq(id)).Take()
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}

				// TODO(vgalaxy): move out tx scope
				res, err := convert.CommentConverterORM(ctx, q, c, redis.InvalidUserId) // 不允许自己关注自己
				if err != nil {
					return err
				}

				resp = &comment.CommentActionResponse{
					Comment: res,
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
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database delete error, nonexistent comment")
				}

				v := tx.Video
				result, err = v.WithContext(ctx).Where(v.ID.Eq(req.VideoId)).Update(v.CommentCount, v.CommentCount.Sub(1))
				if err != nil {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
				}
				if result.RowsAffected != 1 {
					return kerrors.NewBizStatusError(int32(errno.ErrDatabase), "database update error")
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
	// check user
	_, err = pack.QueryUser(ctx, req.UserViewId)
	if err != nil {
		return nil, err
	}

	// check video
	_, err = pack.QueryVideo(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	// do action
	q := query.Use(mysql.DB)
	var comments []*model.Comment
	err = q.Transaction(func(tx *query.Query) error {
		comments, err = tx.Comment.WithContext(ctx).
			Preload(tx.Comment.User).
			Where(tx.Comment.VideoID.Eq(req.VideoId)).
			Find()
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	list := make([]*rpcmodel.Comment, len(comments))
	for i, c := range comments {
		list[i], err = convert.CommentConverterORM(ctx, q, c, req.UserViewId)
		if err != nil {
			return nil, err
		}
	}
	resp = &comment.CommentListResponse{
		CommentList: list,
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}
