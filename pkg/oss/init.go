package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
)

var (
	client      *oss.Client
	videoBucket *oss.Bucket
	coverBucket *oss.Bucket
)

func Init() {
	v := conf.Init()

	c, err := oss.New(v.GetString("aliyun-oss.endpoint"),
		v.GetString("aliyun-oss.ak"),
		v.GetString("aliyun-oss.sk"))
	if err != nil {
		panic(err)
	}

	vb, err := c.Bucket(v.GetString("aliyun-oss.video-bucket-name"))
	if err != nil {
		panic(err)
	}

	cb, err := c.Bucket(v.GetString("aliyun-oss.cover-bucket-name"))
	if err != nil {
		panic(err)
	}

	client = c
	videoBucket = vb
	coverBucket = cb
}
