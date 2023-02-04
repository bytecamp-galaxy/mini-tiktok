package test

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func newExpect(t *testing.T) *httpexpect.Expect {
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,
		BaseURL:  fmt.Sprintf("http://127.0.0.1:8080"), // TODO(vgalaxy): config
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			// httpexpect.NewDebugPrinter(t, false),
			httpexpect.NewCompactPrinter(t),
		},
	})
}

func userRegisterAndPublish(username string, e *httpexpect.Expect) (int64, string) {
	password := utils.RandStringBytesMaskImprSrcUnsafe(15)
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", username).WithQuery("password", password).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	registerResp.Value("status_code").Number().Equal(0)

	userId := int64(registerResp.Value("user_id").Number().Raw())
	token := registerResp.Value("token").String().Raw()

	publishResp := e.POST("/douyin/publish/action/").
		WithMultipart().
		WithFile("data", "../assets/test.mp4").
		WithFormField("token", token).
		WithFormField("title", "test video").
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	publishResp.Value("status_code").Number().Equal(0)

	return userId, token
}
