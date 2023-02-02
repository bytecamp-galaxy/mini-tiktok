package test

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/gavv/httpexpect/v2"
	"net/http"
	"testing"
)

var testUserA = "XK1W5EQRyMdt76C"
var testUserB = "JPgQwIiILqHbZxL"

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

func getTestUserToken(user string, e *httpexpect.Expect) (int64, string) {
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", user).WithQuery("password", user).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := int64(registerResp.Value("user_id").Number().Raw())
	token := registerResp.Value("token").String().Raw()
	return userId, token
}
