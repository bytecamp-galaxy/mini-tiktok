package handler

import (
	"bytes"
	"context"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/oss"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/snowflake"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/google/uuid"
	"net/http"
)

// PublishServiceImpl implements the last service interface defined in the IDL.
type PublishServiceImpl struct{}

// PublishVideo implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishVideo(ctx context.Context, req *publish.PublishRequest) (resp *publish.PublishResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videoData := req.Data
	authorId := req.UserId
	videoTitle := req.Title

	// byte[] -> reader
	reader := bytes.NewReader(videoData)
	// check video type
	filetype := http.DetectContentType(videoData)
	if filetype != "video/mp4" {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidVideoType), err.Error())
	}

	videoUid := uuid.New()
	fileName := videoUid.String() + "." + "mp4"

	v := conf.Init()
	videoBucketName := v.GetString("minio.video-bucket-name")
	//coverBucketName := v.GetString("minio.cover-bucket-name")

	// 上传视频
	//err = minio.UploadFile(videoBucketName, fileName, reader, int64(len(videoData)))
	//if err != nil {
	//	return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	//}

	err = oss.UploadFile(videoBucketName, fileName, reader)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrOSS), err.Error())
	}

	// 获取视频链接
	//playUrl, err := minio.GetFileUrl(videoBucketName, fileName, 0)
	//if err != nil {
	//	return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	//}

	playUrl := oss.GetFileUrl(videoBucketName, fileName)

	// 获取封面
	//coverUid := uuid.New()
	//coverPath := coverUid.String() + "." + "jpg"
	//coverData, err := utils.ReadFrameAsJpeg(playUrl.String())
	//if err != nil {
	//	return nil, kerrors.NewBizStatusError(int32(errno.ErrEncodingFailed), err.Error())
	//}

	// 上传封面
	//coverReader := bytes.NewReader(coverData)
	//err = minio.UploadFile(coverBucketName, coverPath, coverReader, int64(len(coverData)))
	//if err != nil {
	//	return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	//}

	// 获取封面链接
	//coverUrl, err := minio.GetFileUrl(coverBucketName, coverPath, 0)
	//if err != nil {
	//	return nil, kerrors.NewBizStatusError(int32(errno.ErrMinio), err.Error())
	//}

	coverUrl := oss.VideoSnapshot(playUrl)

	// 封装 video
	vid := snowflake.Generate()
	video := &model.Video{
		ID:       vid,
		AuthorID: authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    videoTitle,
	}

	// 保存
	err = query.Video.WithContext(ctx).Create(video)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// load to redis bloom filter
	err = redis.VideoIdAddBF(ctx, vid)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrRedis), err.Error())
	}

	// response to publish server
	resp = &publish.PublishResponse{}
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	// check user
	_, err = pack.QueryUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	_, err = pack.QueryUser(ctx, req.UserViewId)
	if err != nil {
		return nil, err
	}

	// query videos in db
	v := query.Video

	videos, err := v.WithContext(ctx).Preload(v.Author).Order(v.CreatedAt.Desc()).Where(v.AuthorID.Eq(req.UserId)).Find()
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	respVideos := make([]*rpcmodel.Video, len(videos))
	for i, video := range videos {
		respVideos[i], err = convert.VideoConverterORM(ctx, query.Q, video, req.UserViewId)
		if err != nil {
			return nil, err
		}
	}

	resp = &publish.PublishListResponse{
		VideoList: respVideos,
	}
	return resp, nil
}
