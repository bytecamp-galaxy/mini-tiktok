package test

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

func newExpect(t *testing.T) *httpexpect.Expect {
	v := conf.Init()
	return httpexpect.WithConfig(httpexpect.Config{
		Client:   http.DefaultClient,
		BaseURL:  fmt.Sprintf("http://%s:%d", v.GetString("api-server.host"), v.GetInt("api-server.port")),
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}

func userRegister(username string, e *httpexpect.Expect) (int64, string) {
	password := utils.RandStringBytesMaskImprSrcUnsafe(15)
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", username).WithQuery("password", password).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := int64(registerResp.Value("user_id").Number().Raw())
	token := registerResp.Value("token").String().Raw()
	return userId, token
}
