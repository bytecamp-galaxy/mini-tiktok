package pack

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/marmotedu/errors"
)

const (
	SuccessStatusMessage         = "success"
	BrokenInvariantStatusMessage = "broken invariant"
)

type Response struct {
	StatusCode int     `json:"status_code,required"`
	StatusMsg  *string `json:"status_msg,required"`
}

func Error(c *app.RequestContext, err error) {
	e := errors.ParseCoder(err)
	c.JSON(e.HTTPStatus(), &Response{
		StatusCode: e.Code(),
		StatusMsg:  utils.String(fmt.Sprintf("%#+v", err)),
		// output full error stack details, useful for debugging with JSON formatted output
	})
}
