package pack

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/marmotedu/errors"
)

const (
	SuccessStatusMessage = "success"
)

type Response struct {
	StatusCode int     `json:"status_code,required"`
	StatusMsg  *string `json:"status_msg,required"`
}

func Error(c *app.RequestContext, err error) {
	// Output full error stack details, useful for debugging.
	hlog.Errorf("%+v", err)
	e := errors.ParseCoder(err)
	c.JSON(e.HTTPStatus(), &Response{
		StatusCode: e.Code(),
		// Returns the user-safe error string mapped to the error code or the error message if none is specified.
		StatusMsg: utils.String(fmt.Sprintf("%s", err)),
	})
}
