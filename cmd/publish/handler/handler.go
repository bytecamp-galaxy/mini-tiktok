package handler

import (
	"bytes"
	"context"
	"errors"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/convert"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/model"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/mysql"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/dal/query"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/oss"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/pack"
	"github.com/bytecamp-galaxy/mini-tiktok/internal/redis"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/publish"
	"github.com/bytecamp-galaxy/mini-tiktok/kitex_gen/rpcmodel"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/conf"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/errno"
	"github.com/bytecamp-galaxy/mini-tiktok/pkg/utils"
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
	videoReader := bytes.NewReader(videoData)
	// check video type
	filetype := http.DetectContentType(videoData)
	if filetype != "video/mp4" {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrInvalidVideoType), err.Error())
	}

	videoUid := uuid.New()
	videoName := videoUid.String() + "." + "mp4"

	v := conf.Init()
	videoBucketName := v.GetString("oss.video-bucket-name")
	coverBucketName := v.GetString("oss.cover-bucket-name")

	// 上传视频
	err = oss.Upload(videoBucketName, videoName, videoReader, int64(len(videoData)))
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrOss), err.Error())
	}

	// 获取视频链接
	playUrl, err := oss.GetUrl(videoBucketName, videoName)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrOss), err.Error())
	}

	// 获取封面
	coverUid := uuid.New()
	coverName := coverUid.String() + "." + "jpg"
	coverData, err := utils.ReadFrameAsJpeg(playUrl)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrEncodingFailed), err.Error())
	}

	// 上传封面
	coverReader := bytes.NewReader(coverData)
	err = oss.Upload(coverBucketName, coverName, coverReader, int64(len(coverData)))
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrOss), err.Error())
	}

	// 获取封面链接
	coverUrl, err := oss.GetUrl(coverBucketName, coverName)
	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrOss), err.Error())
	}

	// 封装 video
	video := &model.Video{
		AuthorID: authorId,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    videoTitle,
	}

	// db transaction
	q := query.Use(mysql.DB)
	err = q.Transaction(func(tx *query.Query) error {
		// create video
		err = query.Video.WithContext(ctx).Create(video)
		if err != nil {
			return err
		}

		// update user info
		result, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(authorId)).
			Update(query.User.WorkCount, query.User.WorkCount.Add(1))
		if err != nil {
			return err
		}
		if result.RowsAffected != 1 {
			return errors.New("database update error")
		}

		return nil
	})

	if err != nil {
		return nil, kerrors.NewBizStatusError(int32(errno.ErrDatabase), err.Error())
	}

	// redis transaction
	err = func() error {
		// load vid to redis bloom filter
		err = redis.VideoIdAddBF(ctx, video.ID)
		if err != nil {
			return err
		}

		// update redis user info if exists
		exist, err := redis.UserInfoExists(ctx, authorId)
		if err != nil {
			return err
		}
		if exist {
			user, err := redis.UserInfoGet(ctx, authorId)
			if err != nil {
				return err
			}
			user.WorkCount += 1
			err = redis.UserInfoSet(ctx, user)
			if err != nil {
				return err
			}
		}

		return nil
	}()

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
