package minio

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient *minio.Client
)

// Minio 对象存储初始化
func init() {
	v := conf.Init()
	endpoint := fmt.Sprintf("%s:%d", v.GetString("minio.host"), v.GetInt("minio.port"))

	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(v.GetString("minio.ak"), v.GetString("minio.sk"), ""),
	})
	if err != nil {
		panic(err)
	}

	minioClient = client

	err = CreateBucket(v.GetString("minio.video-bucket-name"))
	if err != nil {
		panic(err)
	}
	err = CreateBucket(v.GetString("minio.cover-bucket-name"))
	if err != nil {
		panic(err)
	}
}
