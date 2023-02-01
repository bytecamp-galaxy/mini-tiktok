package handler

import (
	"context"
	favorite "github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionRequest) (resp *favorite.FavoriteActionResponse, err error) {

	switch req.ActionType {
	case 1:
		{
			return BuildFavoriteActionResp(nil), Favorite(ctx, req.UserId, req.VideoId)
		}
	case 2:
		{
			return BuildFavoriteActionResp(nil), DisFavorite(ctx, req.UserId, req.VideoId)
		}

	default:
		panic("FavoriteActionType Error!")
	}
	return

}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteListRequest) (resp *favorite.FavoriteListResponse, err error) {
	videos, err := s._FavoriteList(ctx, req)
	resp = BuildFavoriteListResp(nil)
	resp.VidoeList = videos
	return resp, nil

}

func BuildFavoriteActionResp(err error) *favorite.FavoriteActionResponse {
	if err == nil {
		return FavoriteActionResp()
	}
	return FavoriteActionResp()
}

func FavoriteActionResp() *favorite.FavoriteActionResponse {
	return &favorite.FavoriteActionResponse{StatusCode: int32(0), StatusMsg: nil}
}

func BuildFavoriteListResp(err error) *favorite.FavoriteListResponse {
	if err == nil {
		return favoriteListResp()
	}
	return favoriteListResp()
}

func favoriteListResp() *favorite.FavoriteListResponse {
	return &favorite.FavoriteListResponse{StatusCode: int32(0), StatusMsg: nil}
}
