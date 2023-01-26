package main

import (
	"context"
	feed "github.com/bytecamp-galaxy/mini-tiktok/feed-server/kitex_gen/feed"
)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// GetFeed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) GetFeed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	// TODO: Your code here...
	return
}
