package minio

import (
	"context"
	"io"
	"net/url"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
)

// CreateBucket 创建桶
func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		klog.Error("bucketName invalid")
	}

	location := "beijing"
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			klog.Debugf("bucket %s already exists", bucketName)
			return nil
		} else {
			klog.Debugf("%s", errBucketExists.Error())
			return err
		}
	} else {
		klog.Infof("bucket %s create successfully", bucketName)
	}
	return nil
}

// UploadLocalFile 上传本地文件（提供文件路径）至 minio
func UploadLocalFile(bucketName string, objectName string, filePath string, contentType string) (int64, error) {
	ctx := context.Background()
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		klog.Errorf("local file upload failed, %s", err)
		return 0, err
	}
	klog.Infof("upload %s of size %d successfully", objectName, info.Size)
	return info.Size, nil
}

// UploadFile 上传文件（提供 reader）至 minio
func UploadFile(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	ctx := context.Background()
	n, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectsize, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		klog.Errorf("upload %s of size %d failed, %s", bucketName, objectsize, err)
		return err
	}
	klog.Infof("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileUrl 从 minio 获取文件 Url
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		klog.Errorf("get url of file %s from bucket %s failed, %s", fileName, bucketName, err)
		return nil, err
	}
	return presignedUrl, nil
}

// RemoveOneFile 从 minio 中删除文件
func RemoveOneFile(bucketName string, fileName string) error {
	// 删除一个文件
	ctx := context.Background()
	_ = minioClient.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{GovernanceBypass: true})
	return nil
}
