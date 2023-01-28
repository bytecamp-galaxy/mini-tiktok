package minio

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

const (
	TestBucketName = "tiktoktest"
)

// MinioServerUrl 不能包含 '/' 等特殊字符
// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	CreateBucket(TestBucketName)
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile(TestBucketName, "test.mp4", "./test.mp4", "video/mp4")
	fmt.Println(info, err)
}

func TestUploadFile(t *testing.T) {
	file, _ := os.Open("./test.mp4")
	defer file.Close()
	fi, _ := os.Stat("./test.mp4")
	err := UploadFile(TestBucketName, "testing.mp4", file, fi.Size())
	fmt.Println(err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := GetFileUrl(TestBucketName, "test.mp4", 0)
	fmt.Println(url, err, strings.Split(url.String(), "?")[0])
	fmt.Println(url.Path, url.RawPath)
}

func TestRemoveOneFile(t *testing.T) {
	err := RemoveOneFile(TestBucketName, "ceshi2")
	fmt.Println(err)
}
