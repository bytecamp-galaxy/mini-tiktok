package test

import (
	"flag"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/gavv/httpexpect/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gmeasure"
	"net/http"
	"testing"
	"time"
)

var (
	e *httpexpect.Expect

	token string
	uid   int64

	username = utils.RandStringBytesMaskImprSrcUnsafe(15)
	password = utils.RandStringBytesMaskImprSrcUnsafe(15)

	n    = flag.Int("n", 10, "the maximum number of samples to record")
	path = flag.String("path", "../assets/test.mp4", "path to video")
)

func TestPerformance(t *testing.T) {
	e = newExpect(t)
	RegisterFailHandler(Fail)
	RunSpecs(t, "performance test")
}

// TODO(vgalaxy): use gomega expect for http status code
var _ = Describe("performance test", Ordered, func() {
	It("ping", func() {
		resp := e.GET("/ping").
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		Expect(resp.Value("message").String().Raw()).To(Equal("pong"))
	})

	It("user register", func() {
		resp := e.POST("/douyin/user/register/").
			WithQuery("username", username).WithQuery("password", password).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		Expect(resp.Value("status_code").Number().Raw()).To(Equal(float64(0)))

		token = resp.Value("token").String().Raw()
		uid = int64(resp.Value("user_id").Number().Raw())
	})

	It("user publish", func() {
		experiment := gmeasure.NewExperiment("publish performance")
		AddReportEntry(experiment.Name, experiment)
		experiment.Sample(func(idx int) {
			experiment.MeasureDuration("publish", func() {
				title := utils.RandStringBytesMaskImprSrcUnsafe(15)
				resp := e.POST("/douyin/publish/action/").
					WithMultipart().
					WithFile("data", *path).
					WithFormField("token", token).
					WithFormField("title", title).
					Expect().
					Status(http.StatusOK).
					JSON().Object()
				Expect(resp.Value("status_code").Number().Raw()).To(Equal(float64(0)))
			})
		}, gmeasure.SamplingConfig{N: *n, Duration: time.Minute})
	})

	It("query published videos", func() {
		resp := e.GET("/douyin/publish/list/").
			WithQuery("user_id", uid).WithQuery("token", token).
			Expect().
			Status(http.StatusOK).
			JSON().Object()
		Expect(resp.Value("status_code").Number().Raw()).To(Equal(float64(0)))
		Expect(resp.Value("video_list").Array().Length().Raw()).To(Equal(float64(*n)))
	})
})
