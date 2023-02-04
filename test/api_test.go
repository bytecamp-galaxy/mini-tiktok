package test

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/log"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/gavv/httpexpect/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
	"net/http"
	"testing"
)

var logger *zap.Logger
var e *httpexpect.Expect

func TestAPI(t *testing.T) {
	logger = log.GetTestLogger()
	e = newExpect(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "API TESTS")
}

// TODO(vgalaxy): use gomega expect, now test always success
var _ = Describe("API TESTS", Ordered, func() {
	It("ping", func() {
		e.GET("/ping").Expect().Status(http.StatusOK).JSON().Object().Value("message").Equal("pong")
	})

	usernameA := utils.RandStringBytesMaskImprSrcUnsafe(15)
	passwordA := utils.RandStringBytesMaskImprSrcUnsafe(15)
	videoTitleA := "user A video"
	commentAA := "user A comment user A video"
	commentAB := "user A comment user B video"

	var tokenA string
	var userIdA int64
	var videoIdA int64
	var commentIdAA int64
	var commentIdAB int64

	usernameB := utils.RandStringBytesMaskImprSrcUnsafe(15)
	passwordB := utils.RandStringBytesMaskImprSrcUnsafe(15)
	videoTitleB := "user B video"
	commentBA := "user B comment user A video"
	commentBB := "user B comment user B video"

	var tokenB string
	var userIdB int64
	var videoIdB int64
	var commentIdBA int64
	var commentIdBB int64

	utils.Use(usernameA, usernameB, passwordA, passwordB, videoTitleA, videoTitleB, commentAA, commentAB, commentBA, commentBB)
	utils.Use(tokenA, tokenB, userIdA, userIdB, videoIdA, videoIdB, commentIdAA, commentIdAB, commentIdBA, commentIdBB)

	It("user A register", func() {
		resp := e.POST("/douyin/user/register/").
			WithQuery("username", usernameA).WithQuery("password", passwordA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		tokenA = resp.Value("token").String().Raw()
		userIdA = int64(resp.Value("user_id").Number().Raw())
	})

	It("user B register", func() {
		resp := e.POST("/douyin/user/register/").
			WithQuery("username", usernameB).WithQuery("password", passwordB).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		tokenB = resp.Value("token").String().Raw()
		userIdB = int64(resp.Value("user_id").Number().Raw())
	})

	It("user A login", func() {
		resp := e.POST("/douyin/user/login/").
			WithQuery("username", usernameA).WithQuery("password", passwordA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		resp.Value("user_id").Number().Equal(userIdA)

		tokenA = resp.Value("token").String().Raw()
	})

	It("user A query user A info", func() {
		resp := e.GET("/douyin/user/").
			WithQuery("user_id", userIdA).WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		user := resp.Value("user").Object()
		user.NotEmpty()
		user.Value("id").Number().Equal(userIdA)
		user.Value("name").String().Equal(usernameA)
		user.Value("follow_count").Number().Equal(0)
		user.Value("follower_count").Number().Equal(0)
		user.Value("is_follow").Boolean().Equal(false)
	})

	It("user A query user B info", func() {
		resp := e.GET("/douyin/user/").
			WithQuery("user_id", userIdB).WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		user := resp.Value("user").Object()
		user.NotEmpty()
		user.Value("id").Number().Equal(userIdB)
		user.Value("name").String().Equal(usernameB)
		user.Value("follow_count").Number().Equal(0)
		user.Value("follower_count").Number().Equal(0)
		user.Value("is_follow").Boolean().Equal(false)
	})

	// TODO: follow action

	It("user A publish video", func() {
		resp := e.POST("/douyin/publish/action/").
			WithMultipart().
			WithFile("data", "../assets/test.mp4").
			WithFormField("token", tokenA).
			WithFormField("title", videoTitleA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		// TODO: wait for publish
		// time.Sleep(5 * time.Second)
	})

	It("user A query user A published videos", func() {
		resp := e.GET("/douyin/publish/list/").
			WithQuery("user_id", userIdA).WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		videoIdA = int64(video.Value("id").Number().Raw())

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(0)
		video.Value("comment_count").Number().Equal(0)
		video.Value("is_favorite").Boolean().Equal(false)
	})

	It("user A query user B published videos", func() {
		resp := e.GET("/douyin/publish/list/").
			WithQuery("user_id", userIdB).WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		resp.Value("video_list").Array().Length().Equal(0)
	})

	It("user B favorite user A video", func() {
		resp := e.POST("/douyin/favorite/action/").
			WithQuery("token", tokenB).WithQuery("video_id", videoIdA).WithQuery("action_type", 1).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
	})

	It("user A query user A published videos", func() {
		resp := e.GET("/douyin/publish/list/").
			WithQuery("user_id", userIdA).WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		videoIdA = int64(video.Value("id").Number().Raw())

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(1) // add one
		video.Value("comment_count").Number().Equal(0)
		video.Value("is_favorite").Boolean().Equal(false)
	})

	It("user B query user B favorite videos", func() {
		resp := e.GET("/douyin/favorite/list/").
			WithQuery("token", tokenB).WithQuery("user_id", userIdB).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		video.Value("id").Number().Equal(videoIdA)

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(1) // add one
		video.Value("comment_count").Number().Equal(0)
		video.Value("is_favorite").Boolean().Equal(true) // note
	})

	It("user B comment user A video", func() {
		resp := e.POST("/douyin/comment/action/").
			WithQuery("token", tokenB).
			WithQuery("video_id", videoIdA).
			WithQuery("action_type", 1).
			WithQuery("comment_text", commentBA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		comment := resp.Value("comment").Object()
		commentIdBA = int64(comment.Value("id").Number().Raw())

		comment.Value("content").String().Equal(commentBA)
		comment.Value("create_date").String().NotEmpty()

		user := comment.Value("user").Object()
		user.Value("id").Number().Equal(userIdB)
		user.Value("name").String().Equal(usernameB)
		user.Value("follow_count").Number().Equal(0)
		user.Value("follower_count").Number().Equal(0)
		user.Value("is_follow").Boolean().Equal(false)
	})

	It("user B query user A published videos", func() {
		resp := e.GET("/douyin/publish/list/").
			WithQuery("user_id", userIdA).WithQuery("token", tokenB).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		videoIdA = int64(video.Value("id").Number().Raw())

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(1)
		video.Value("comment_count").Number().Equal(1) // add one
		video.Value("is_favorite").Boolean().Equal(false)
	})

	It("user A query user A video comments", func() {
		resp := e.GET("/douyin/comment/list/").
			WithQuery("token", tokenA).WithQuery("video_id", videoIdA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		resp.Value("comment_list").Array().Length().Equal(1)

		comment := resp.Value("comment_list").Array().First().Object()
		comment.Value("id").Number().Equal(commentIdBA)

		comment.Value("content").String().Equal(commentBA)
		comment.Value("create_date").String().NotEmpty()

		user := comment.Value("user").Object()
		user.Value("id").Number().Equal(userIdB)
		user.Value("name").String().Equal(usernameB)
		user.Value("follow_count").Number().Equal(0)
		user.Value("follower_count").Number().Equal(0)
		user.Value("is_follow").Boolean().Equal(false)
	})

	It("user A feed", func() {
		resp := e.GET("/douyin/feed/").
			WithQuery("token", tokenA).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		nextTime := int64(resp.Value("next_time").Number().Raw()) // `next_time` required here

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		video.Value("id").Number().Equal(videoIdA)

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(1)
		video.Value("comment_count").Number().Equal(1)
		video.Value("is_favorite").Boolean().Equal(false)

		resp = e.GET("/douyin/feed/").
			WithQuery("token", tokenA).WithQuery("latest_time", nextTime).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		resp.Value("video_list").Array().Empty() // empty
	})

	It("user B feed", func() {
		resp := e.GET("/douyin/feed/").
			WithQuery("token", tokenB).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		nextTime := int64(resp.Value("next_time").Number().Raw()) // `next_time` required here

		resp.Value("video_list").Array().Length().Equal(1)
		video := resp.Value("video_list").Array().First().Object()
		video.Value("id").Number().Equal(videoIdA)

		author := video.Value("author").Object()
		author.Value("id").Number().Equal(userIdA)
		author.Value("name").String().Equal(usernameA)
		author.Value("follow_count").Number().Equal(0)
		author.Value("follower_count").Number().Equal(0)
		author.Value("is_follow").Boolean().Equal(false)

		video.Value("play_url").String().NotEmpty()
		video.Value("cover_url").String().NotEmpty()

		video.Value("title").String().Equal(videoTitleA)
		video.Value("favorite_count").Number().Equal(1)
		video.Value("comment_count").Number().Equal(1)
		video.Value("is_favorite").Boolean().Equal(true) // favorite

		resp = e.GET("/douyin/feed/").
			WithQuery("token", tokenB).WithQuery("latest_time", nextTime).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		resp.Value("status_code").Number().Equal(0)
		resp.Value("video_list").Array().Empty() // empty
	})
})
