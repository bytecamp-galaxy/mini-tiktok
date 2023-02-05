package handler

import (
	"context"
	favorite "github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {
	resp = &favorite.FavoriteActionResponse{}
	switch req.ActionType {
	case 1:
		{
			e := doFavorite(ctx, req.UserId, req.VideoId)
			if e != nil {
				return nil, e
			}
			return resp, nil
		}
	case 2:
		{
			e := doUnfavorite(ctx, req.UserId, req.VideoId)
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
	videos, err := s.favoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	resp = &favorite.FavoriteListResponse{
		VideoList: videos,
	}
	return resp, nil
}
