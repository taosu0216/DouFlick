package service

import (
	"context"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"go.uber.org/zap"
	"strconv"
	"usersvr/log"
	"videosvr/init/minio"
	"videosvr/utils"
)

/*
  // 获取视频列表
  rpc GetPublishVideoList(GetPublishVideoListRequest) returns (GetPublishVideoListResponse);
  rpc PublishVideo(PublishVideoRequest) returns (PublishVideoResponse);
  rpc GetFeedList(GetFeedListRequest) returns (GetFeedListResponse);
  rpc GetVideoInfoList(GetVideoInfoListReq) returns (GetVideoInfoListRsp);
  rpc GetFavoriteVideoList(GetFavoriteVideoListReq) returns (GetFavoriteVideoListRsp);
  // 更新这个视频的获赞数
  rpc UpdateFavoriteCount(UpdateFavoriteCountReq) returns (UpdateFavoriteCountRsp);
  // 更新这个视频的评论数
  rpc UpdateCommentCount(UpdateCommentCountReq) returns (UpdateCommentCountRsp);
*/

type VideoService struct {
	pb.VideoServiceClient
}

func (VideoService) PublishVideo(ctx context.Context, req *pb.PublishVideoRequest) (*pb.PublishVideoResponse, error) {
	minioClient := minio.GetMinio()
	videoUrl, err := minioClient.UploadFile("video", req.SaveFile, strconv.FormatInt(req.UserId, 10))
	if err != nil {
		log.Error("UploadFile err:%v", err)
		return nil, err
	}
	//生成视频封面(截取视频第一帧)
	imageFile, err := utils.GetImage(req.SaveFile)
	if err != nil {
		log.Error("GetImage err:%v", err)
		return nil, err
	}

	picUrl, err := minioClient.UploadFile("pic", imageFile, strconv.FormatInt(req.UserId, 10))
	if err != nil {
		log.Error("UploadFile err", zap.Error(err))
		picUrl = "https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/7909abe413ec4a1e82032d2beb810157~tplv-k3u1fbpfcp-zoom-in-crop-mark:1304:0:0:0.awebp?"
	}
	//TODO: 存储到db未实现
	return &pb.PublishVideoResponse{}, nil
}
