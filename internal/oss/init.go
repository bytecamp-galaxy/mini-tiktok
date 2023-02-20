package oss

import (
	"errors"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/oss/aliyun"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/oss/local"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/oss/minio"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"io"
)

var (
	AliyunInstance = "aliyun"
	MinioInstance  = "minio"
	LocalInstance  = "local"
)

var instance Instance = minio.OssMinio{}

type Instance interface {
	// Upload 上传文件
	Upload(bucketName string, objectName string, reader io.Reader, objectsize int64) error
	// GetUrl 获取文件 url
	GetUrl(bucketName string, objectName string) (string, error)
	// CreateBucket 创建桶
	CreateBucket(bucketName string) error
}

func Upload(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	return instance.Upload(bucketName, objectName, reader, objectsize)
}

func GetUrl(bucketName string, objectName string) (string, error) {
	return instance.GetUrl(bucketName, objectName)
}

func CreateBucket(bucketName string) error {
	return instance.CreateBucket(bucketName)
}

// Init only called by publish service
func Init(ossType string) {
	switch ossType {
	case AliyunInstance:
		aliyun.Init()
		instance = aliyun.OssAliyun{}
	case MinioInstance:
		minio.Init()
		instance = minio.OssMinio{}
	case LocalInstance:
		local.Init()
		instance = local.OssLocal{}
	default:
		panic(errors.New("oss type unsupported"))
	}

	v := conf.Init()
	if v.GetBool("oss.init-bucket") {
		err := CreateBucket(v.GetString("oss.video-bucket-name"))
		if err != nil {
			panic(err)
		}
		err = CreateBucket(v.GetString("oss.cover-bucket-name"))
		if err != nil {
			panic(err)
		}
	}
}
