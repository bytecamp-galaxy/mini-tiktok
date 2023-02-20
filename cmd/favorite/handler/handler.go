package handler

import (
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	favorite "github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	// check video
	v, err := pack.QueryVideo(ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	// do action
	resp = &favorite.FavoriteActionResponse{}
	switch req.ActionType {
	case 1:
		{
			e := doFavorite(ctx, req.UserId, req.VideoId, v.AuthorID)
			if e != nil {
				return nil, e
			}
			return resp, nil
		}
	case 2:
		{
			e := doUnfavorite(ctx, req.UserId, req.VideoId, v.AuthorID)
			if e != nil {
				return nil, e
			}
			return resp, nil
		}
	default:
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), "unknown action type")
	}
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	_, err = pack.QueryUser(ctx, req.UserViewId)
	if err != nil {
		return nil, err
	}

	fr := query.FavoriteRelation
	rs, err := fr.WithContext(ctx).Preload(fr.Video).Preload(fr.Video.Author).Where(fr.UserID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	videos := make([]*rpcmodel.Video, len(rs))
	for i, r := range rs {
		videos[i], err = convert.VideoConverterORM(ctx, query.Q, &r.Video, req.UserViewId)
		if err != nil {
			return nil, err
		}
	}

	resp = &favorite.FavoriteListResponse{
		VideoList: videos,
	}
	return resp, nil
}
