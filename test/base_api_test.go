package test

import (
	"bytes"
	"github.com/bytecamp-galaxy/mini-tiktok/api-server/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"testing"
)

func TestPing(t *testing.T) {
	h := server.Default()
	h.GET("/ping", handler.Ping)
	w := ut.PerformRequest(h.Engine, "GET", "/ping", &ut.Body{bytes.NewBufferString("1"), 1},
		ut.Header{"Connection", "close"})
	resp := w.Result()
	assert.DeepEqual(t, 200, resp.StatusCode())
	assert.DeepEqual(t, "{\"message\":\"pong\"}", string(resp.Body()))
}
