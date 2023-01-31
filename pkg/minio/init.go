package minio

import (
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/constants"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	minioClient     *minio.Client
	Endpoint        = constants.MinioServerUrl
	AccessKeyId     = constants.MinioAccessKey
	SecretAccessKey = constants.MinioSecretKey
	UseSSL          = constants.MinioUseSSL
)

// Minio 对象存储初始化
func init() {
	client, err := minio.New(Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(AccessKeyId, SecretAccessKey, ""),
		Secure: UseSSL,
	})
	if err != nil {
		klog.Errorf("minio client init failed: %v", err)
	}
	// fmt.Println(client)
	klog.Debug("minio client init successfully")
	minioClient = client
}
