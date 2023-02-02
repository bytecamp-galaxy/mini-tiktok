package handler

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bytecamp-galaxy/mini-tiktok/cmd/api/biz/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/minio"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"image"
	"image/jpeg"
	"os"
	"strings"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishRequest) (resp *publish.PublishResponse, err error) {
	videoData := req.Data
	authorId := req.UserId
	videoTitle := req.Title
	// // 获取后缀
	// filetype := http.DetectContentType(videoData)

	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	videoUid := uuid.New()
	fileName := videoUid.String() + "." + "mp4"

	v := conf.Init()
	videoBucketName := v.GetString("minio.video-bucket-name")
	coverBucketName := v.GetString("minio.cover-bucket-name")

	// 上传视频
	err = minio.UploadFile(videoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// 获取视频链接
	url, err := minio.GetFileUrl(videoBucketName, fileName, 0)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// 获取封面
	coverUid := uuid.New()
	coverPath := coverUid.String() + "." + "jpg"
	coverData, err := readFrameAsJpeg(playUrl)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(coverBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(coverBucketName, coverPath, 0)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// CoverUrl := strings.Split(coverUrl.String(), "?")[0]
	// 获取 user
	author, err := getAuthorInfo(ctx, authorId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrUnknown), err.Error())
	}

	// 封装 video
	video := &model.Video{
		AuthorID: authorId,
		Author:   *author,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl.String(),
		Title:    videoTitle,
	}

	// 保存
	err = query.Video.WithContext(ctx).Create(video)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// response to publish server
	resp = &publish.PublishResponse{}
	return resp, nil
}

func getAuthorInfo(ctx context.Context, uid int64) (data *model.User, err error) {
	// set up connection with user server
	cli, err := rpc.InitUserClient()
	if err != nil {
		return nil, err
	}

	// call rpc service
	reqRpc := &user.UserQueryRequest{
		UserId: uid,
	}

	respRpc, err := (*cli).UserQuery(ctx, reqRpc)
	if err != nil {
		// TODO(vgalaxy): use BizStatusErrorIface
		return nil, err
	}

	authorData := respRpc.User
	author := &model.User{
		ID:             authorData.Id,
		Username:       authorData.Name,
		FollowerCount:  authorData.FollowerCount,
		FollowingCount: authorData.FollowCount,
	}
	return author, err
}

// ReadFrameAsJpeg
// 从视频流中截取一帧并返回 需要在本地环境中安装 ffmpeg 并将 bin 添加到环境变量
func readFrameAsJpeg(filePath string) ([]byte, error) {
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
