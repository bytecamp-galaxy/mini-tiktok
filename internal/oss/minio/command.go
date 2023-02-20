package minio

import (
	"context"
	"errors"
	"io"
	"net/url"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

type OssMinio struct {
}

func (oss OssMinio) Upload(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	ctx := context.Background()
	n, err := client.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		klog.Errorf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	klog.Infof("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

func (oss OssMinio) GetUrl(bucketName string, objectName string) (string, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	// TODO: expire time
	expires := time.Second * 60 * 60 * 24
	presignedUrl, err := client.PresignedGetObject(ctx, bucketName, objectName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", objectName, bucketName, err)
		return "", err
	}
	return presignedUrl.String(), nil
}

func (oss OssMinio) CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		return errors.New("bucket name invalid")
	}

	// TODO: bucket location
	location := "beijing"
	ctx := context.Background()
	err := client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		klog.Errorf("create bucket %s failed, %s", bucketName, err)
		return err
	}

	klog.Infof("create bucket %s successfully", bucketName)
	return nil
}
