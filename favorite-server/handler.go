package main

import (
	"context"
	favorite "github.com/bytecamp-galaxy/mini-tiktok/favorite-server/kitex_gen/favorite"
)

// FavoriteServiceImpl implements the last service interface defined in the IDL.
type FavoriteServiceImpl struct{}

// FavoriteAction implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteAction(ctx context.Context, req *favorite.FavoriteActionsRequest) (resp *favorite.FavoriteActionResponse, err error) {
	// TODO: Your code here...

	return
}

// FavoriteList implements the FavoriteServiceImpl interface.
func (s *FavoriteServiceImpl) FavoriteList(ctx context.Context, req *favorite.FavoriteActionsRequest) (resp *favorite.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}
