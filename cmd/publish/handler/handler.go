package handler

import (
	"bytes"
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/rpc"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/user"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/minio"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
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
		return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	}

	// 获取视频链接
	playUrl, err := minio.GetFileUrl(videoBucketName, fileName, 0)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	}

	// 获取封面
	coverUid := uuid.New()
	coverPath := coverUid.String() + "." + "jpg"
	coverData, err := utils.ReadFrameAsJpeg(playUrl.String())
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrEncodingFailed), err.Error())
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = minio.UploadFile(coverBucketName, coverPath, coverReader, int64(len(coverData)))
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	}

	// 获取封面链接
	coverUrl, err := minio.GetFileUrl(coverBucketName, coverPath, 0)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	}

	// 获取 user
	author, err := getAuthorInfo(ctx, authorId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRPCMutualCall), err.Error())
	}

	// 封装 video
	video := &model.Video{
		AuthorID: authorId,
		Author:   *author,
		PlayUrl:  playUrl.String(),
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
