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

func getTestUserToken(user string, e *httpexpect.Expect) (int, string) {
	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", user).WithQuery("password", user).
		WithFormField("username", user).WithFormField("password", user).
		Expect().
		Status(http.StatusOK).
		JSON().Object()

	userId := 0
	token := registerResp.Value("token").String().Raw()
	if len(token) == 0 {
		loginResp := e.POST("/douyin/user/login/").
			WithQuery("username", user).WithQuery("password", user).
			WithFormField("username", user).WithFormField("password", user).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		loginToken := loginResp.Value("token").String()
		loginToken.Length().Gt(0)
		token = loginToken.Raw()
		userId = int(loginResp.Value("user_id").Number().Raw())
	} else {
		userId = int(registerResp.Value("user_id").Number().Raw())
	}
	return userId, token
}
