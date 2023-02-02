package minio

import (
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

var (
	minioClient *minio.Client
)

// Minio 对象存储初始化
func init() {
	v := conf.Init()
	endpoint := fmt.Sprintf("%s:%d", v.GetString("minio.host"), v.GetInt("minio.port"))
	client, err := minio.New(endpoint, &minio.Options{})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	klog.Debug("minio client init successfully")
	minioClient = client
}
