package aliyun

import (
	aliyunoss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
)

var (
	client      *aliyunoss.Client
	videoBucket *aliyunoss.Bucket
	coverBucket *aliyunoss.Bucket
)

func Init() {
	v := conf.Init()

	c, err := aliyunoss.New(v.GetString("oss.endpoint"),
		v.GetString("oss.ak"),
		v.GetString("oss.sk"))
	if err != nil {
		panic(err)
	}

	vb, err := c.Bucket(v.GetString("oss.video-bucket-name"))
	if err != nil {
		panic(err)
	}

	cb, err := c.Bucket(v.GetString("oss.cover-bucket-name"))
	if err != nil {
		panic(err)
	}

	client = c
	videoBucket = vb
	coverBucket = cb
}
