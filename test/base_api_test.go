package test

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"net/http"
	"testing"
)

func TestFeed(t *testing.T) {
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

	for _, element := range feedResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}

func TestUserAction(t *testing.T) {
	e := newExpect(t)

	username := utils.RandStringBytesMaskImprSrcUnsafe(15)
	password := utils.RandStringBytesMaskImprSrcUnsafe(15)

	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", username).WithQuery("password", password).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	registerResp.Value("status_code").Number().Equal(0)

	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", username).WithQuery("password", password).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	loginResp.Value("status_code").Number().Equal(0)

	userId := int64(loginResp.Value("user_id").Number().Raw())
	token := loginResp.Value("token").String().Raw()
	userResp := e.GET("/douyin/user/").
		WithQuery("user_id", userId).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	userResp.Value("status_code").Number().Equal(0)

	userInfo := userResp.Value("user").Object()
	userInfo.NotEmpty()
	userInfo.Value("id").Number().Equal(userId)
	userInfo.Value("name").String().Equal(username)
}

func TestPublish(t *testing.T) {
	e := newExpect(t)

	username := utils.RandStringBytesMaskImprSrcUnsafe(15)
	userId, token := userRegisterAndPublish(username, e)

	publishResp := e.POST("/douyin/publish/action/").
		WithMultipart().
		WithFile("data", "../assets/test.mp4").
		WithFormField("token", token).
		WithFormField("title", "test video").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishResp.Value("status_code").Number().Equal(0)

	publishListResp := e.GET("/douyin/publish/list/").
		WithQuery("user_id", userId).WithQuery("token", token).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishListResp.Value("status_code").Number().Equal(0)
	publishListResp.Value("video_list").Array().Length().Gt(0)

	for _, element := range publishListResp.Value("video_list").Array().Iter() {
		video := element.Object()
		video.ContainsKey("id")
		video.ContainsKey("author")
		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()
	}
}
