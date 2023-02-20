package aliyun

import (
	"errors"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"io"
)

type OssAliyun struct {
}

func (oss OssAliyun) Upload(bucketName string, objectName string, reader io.Reader, objectsize int64) (err error) {
	v := conf.Init()
	switch bucketName {
	case v.GetString("oss.video-bucket-name"):
		err = videoBucket.PutObject(objectName, reader)
	case v.GetString("oss.cover-bucket-name"):
		err = coverBucket.PutObject(objectName, reader)
	default:
		return errors.New("bucket name does not exist")
	}
	return
}

func (oss OssAliyun) GetUrl(bucketName string, objectName string) (string, error) {
	v := conf.Init()
	playUrl := fmt.Sprintf("https://%s.%s/%s",
		v.GetString("oss.video-bucket-name"),
		v.GetString("oss.endpoint"),
		objectName)
	switch bucketName {
	case v.GetString("oss.video-bucket-name"):
		return playUrl, nil
	case v.GetString("oss.cover-bucket-name"):
		// return snapshot(playUrl), nil
		return fmt.Sprintf("https://%s.%s/%s",
			bucketName,
			v.GetString("oss.endpoint"),
			objectName), nil
	default:
		return "", errors.New("bucket name does not exist")
	}
}

func (oss OssAliyun) CreateBucket(bucketName string) error {
	// TODO: implement me
	panic("implement me")
}
