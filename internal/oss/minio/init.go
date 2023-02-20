package minio

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	client *minio.Client
)

func Init() {
	v := conf.Init()
	endpoint := fmt.Sprintf("%s:%d", v.GetString("oss.host"), v.GetInt("oss.port"))

	c, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(v.GetString("oss.ak"), v.GetString("oss.sk"), ""),
	})
	if err != nil {
		panic(err)
	}

	client = c
}
