package oss

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestListBuckets(t *testing.T) {
	Init()

	lsRes, err := client.ListBuckets()
	if err != nil {
		panic(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println(bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {
	v := conf.Init()
	Init()

	f, err := os.Open("assets/test.mp4")
	assert.NoError(t, err)
	defer f.Close()
	err = UploadFile(v.GetString("aliyun-oss.video-bucket-name"), "test.mp4", f)
	assert.NoError(t, err)
}
