package oss

import (
	"errors"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"io"
)

func UploadFile(bucketName string, objectName string, reader io.Reader) (err error) {
	v := conf.Init()
	switch bucketName {
	case v.GetString("aliyun-oss.video-bucket-name"):
		err = videoBucket.PutObject(objectName, reader)
	case v.GetString("aliyun-oss.cover-bucket-name"):
		err = coverBucket.PutObject(objectName, reader)
	default:
		return errors.New("bucket name does not exist")
	}
	return
}

func GetFileUrl(bucketName string, objectName string) string {
	v := conf.Init()
	return fmt.Sprintf("https://%s.%s/%s",
		bucketName,
		v.GetString("endpoint"),
		objectName)
}

func VideoSnapshot(playUrl string) string {
	return fmt.Sprintf("%d?x-oss-process=video/snapshot,t_0,f_jpg,w_0,h_0,m_fast",
		playUrl)
}
