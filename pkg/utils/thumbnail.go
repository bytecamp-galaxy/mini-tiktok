package utils

import (
	"bytes"
	"github.com/bakape/thumbnailer/v2"
	"image/jpeg"
	"io"
)

// GetThumbnail Generate JPEG thumbnail from video
func GetThumbnail(input io.ReadSeeker) ([]byte, error) {
	_, thumb, err := thumbnailer.Process(input, thumbnailer.Options{})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, thumb, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
