package local

import "io"

type OssLocal struct {
}

func (o OssLocal) Upload(bucketName string, objectName string, reader io.Reader, objectsize int64) error {
	// TODO: implement me
	panic("implement me")
}

func (o OssLocal) GetUrl(bucketName string, objectName string) (string, error) {
	// TODO: implement me
	panic("implement me")
}

func (o OssLocal) CreateBucket(bucketName string) error {
	// TODO: implement me
	panic("implement me")
}
