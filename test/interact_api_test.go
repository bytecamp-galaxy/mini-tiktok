package test

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestFavorite(t *testing.T) {
	e := newExpect(t)

	username := utils.RandStringBytesMaskImprSrcUnsafe(15)
	userId, token := userRegisterAndPublish(username, e)

	feedResp := e.GET("/douyin/feed/").
		WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	firstVideo := feedResp.Value("video_list").Array().First().Object()
	videoId := int64(firstVideo.Value("id").Number().Raw())

	favoriteResp := e.POST("/douyin/favorite/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 1).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteResp.Value("status_code").Number().Equal(0)

	favoriteListResp := e.GET("/douyin/favorite/list/").
		WithQuery("token", token).WithQuery("user_id", userId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	favoriteListResp.Value("status_code").Number().Equal(0)

	for _, element := range favoriteListResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}

func TestComment(t *testing.T) {
	e := newExpect(t)

	username := utils.RandStringBytesMaskImprSrcUnsafe(15)
	_, token := userRegisterAndPublish(username, e)

	feedResp := e.GET("/douyin/feed/").
		WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	feedResp.Value("status_code").Number().Equal(0)
	feedResp.Value("video_list").Array().Length().Gt(0)
	firstVideo := feedResp.Value("video_list").Array().First().Object()
	videoId := int64(firstVideo.Value("id").Number().Raw())

	addCommentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 1).WithQuery("comment_text", "test comment").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	addCommentResp.Value("status_code").Number().Equal(0)
	addCommentResp.Value("comment").Object().Value("id").Number().Gt(0)
	addCommentResp.Value("comment").Object().Value("user").Object().Value("id").Number().Gt(0)
	commentId := int64(addCommentResp.Value("comment").Object().Value("id").Number().Raw())

	commentListResp := e.GET("/douyin/comment/list/").
		WithQuery("token", token).WithQuery("video_id", videoId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	commentListResp.Value("status_code").Number().Equal(0)

	containTestComment := false
	for _, element := range commentListResp.Value("comment_list").Array().Iter() {
		comment := element.Object()
		comment.ContainsKey("id")
		comment.ContainsKey("user")
		comment.Value("user").Object().Value("id").NotEqual(0)
		comment.Value("content").String().NotEmpty()
		comment.Value("create_date").String().NotEmpty()
		if int64(comment.Value("id").Number().Raw()) == commentId {
			containTestComment = true
		}
	}
	assert.True(t, containTestComment, "Can't find test comment in list")

	delCommentResp := e.POST("/douyin/comment/action/").
		WithQuery("token", token).WithQuery("video_id", videoId).WithQuery("action_type", 2).WithQuery("comment_id", commentId).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	delCommentResp.Value("status_code").Number().Equal(0)
}
