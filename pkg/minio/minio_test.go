package minio

import (
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
)

const (
	TestBucketName = "tiktoktest"
)

func TestVideoType(t *testing.T) {
	f, err := os.ReadFile("../../assets/test.mp4")
	assert.NoError(t, err)
	filetype := http.DetectContentType(f)
	log.Println(filetype)
}

// MinioServerUrl 不能包含 '/' 等特殊字符
// bucket name 只能用小写字母
func TestCreateBucket(t *testing.T) {
	assert.NoError(t, CreateBucket(TestBucketName))
}

func TestUploadLocalFile(t *testing.T) {
	info, err := UploadLocalFile(TestBucketName, "test.mp4", "assets/test.mp4", "video/mp4")
	assert.NoError(t, err)
	log.Println(info)
}

func TestUploadFile(t *testing.T) {
	f, err := os.Open("assets/test.mp4")
	assert.NoError(t, err)
	defer f.Close()
	info, err := os.Stat("assets/test.mp4")
	assert.NoError(t, err)
	err = UploadFile(TestBucketName, "test.mp4", f, info.Size())
	assert.NoError(t, err)
}

func TestGetFileUrl(t *testing.T) {
	url, err := GetFileUrl(TestBucketName, "test.mp4", 0)
	assert.NoError(t, err)
	log.Println(strings.Split(url.String(), "?")[0], url.Path)
}

func TestRemoveOneFile(t *testing.T) {
	assert.NoError(t, RemoveOneFile(TestBucketName, "test.mp4"))
}
