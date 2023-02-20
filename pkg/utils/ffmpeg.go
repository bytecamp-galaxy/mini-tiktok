package utils

import (
	"bytes"
	"context"
	"fmt"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"gopkg.in/vansante/go-ffprobe.v2"
	"image"
	"image/jpeg"
	"os"

	"io"
)

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装 ffmpeg 并将 bin 添加到环境变量
func ReadFrameAsJpeg(filePath string) ([]byte, error) {
	reader := bytes.NewBuffer(nil)
	err := ffmpeg.Input(filePath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(reader, os.Stdout).
		Run()
	if err != nil {
		return nil, err
	}
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)

	return buf.Bytes(), err
}

func ParseVideoResolution(ctx context.Context, reader io.Reader) (weight int, height int, err error) {
	var data *ffprobe.ProbeData
	data, err = ffprobe.ProbeReader(ctx, reader)
	if err != nil {
		return
	}

	weight = data.Streams[0].Width
	height = data.Streams[0].Height
	return
}
