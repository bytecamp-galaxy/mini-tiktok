package test

import (
	"go.uber.org/zap"
	"net/http"
	"testing"
)

func TestCornerCase(t *testing.T) {
	e := newExpect(t)
	l := zap.NewExample()

	describe := func(text string, fn func()) {
		l.Info(text)
		fn()
	}

	describe("ping", func() {
		e.GET("/ping").Expect().Status(http.StatusOK).JSON().Object().Value("message").Equal("pong")
	})
}
