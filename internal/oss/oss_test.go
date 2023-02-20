package oss

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
)

const (
	testBucketName = "minio-test"
	objectName     = "test.mp4"
	filePath       = "../../assets/test.mp4"
)

func TestVideoType(t *testing.T) {
	f, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	filetype := http.DetectContentType(f)
	assert.Equal(t, filetype, "video/mp4")
}

func TestMinio(t *testing.T) {
	Init(MinioInstance)
	assert.NoError(t, CreateBucket(testBucketName))

	reader, err := os.Open(filePath)
	assert.NoError(t, err)
	defer reader.Close()

	info, err := os.Stat(filePath)
	assert.NoError(t, err)

	err = Upload(testBucketName, objectName, reader, info.Size())
	assert.NoError(t, err)

	url, err := GetUrl(testBucketName, objectName)
	assert.NoError(t, err)
	log.Println(url)
}

func TestAliyun(t *testing.T) {
	Init(AliyunInstance)
	assert.NoError(t, CreateBucket(testBucketName))

	reader, err := os.Open(filePath)
	assert.NoError(t, err)
	defer reader.Close()

	info, err := os.Stat(filePath)
	assert.NoError(t, err)

	err = Upload(testBucketName, objectName, reader, info.Size())
	assert.NoError(t, err)

	url, err := GetUrl(testBucketName, objectName)
	assert.NoError(t, err)
	log.Println(url)
}

func TestLocal(t *testing.T) {
	Init(LocalInstance)
	assert.NoError(t, CreateBucket(testBucketName))

	reader, err := os.Open(filePath)
	assert.NoError(t, err)
	defer reader.Close()

	info, err := os.Stat(filePath)
	assert.NoError(t, err)

	err = Upload(testBucketName, objectName, reader, info.Size())
	assert.NoError(t, err)

	url, err := GetUrl(testBucketName, objectName)
	assert.NoError(t, err)
	log.Println(url)
}
