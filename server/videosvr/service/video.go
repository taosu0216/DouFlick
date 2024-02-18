package service

import (
	"context"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"go.uber.org/zap"
	"strconv"
	"time"
	"videosvr/config"
	"videosvr/init/minio"
	"videosvr/log"
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
	pb.UnimplementedVideoServiceServer
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
	err = utils.InsertVideo(req.UserId, videoUrl, picUrl, req.Title)
	return &pb.PublishVideoResponse{}, nil
}
func (VideoService) UpdateFavoriteCount(ctx context.Context, req *pb.UpdateFavoriteCountReq) (*pb.UpdateFavoriteCountRsp, error) {
	resp := &pb.CommonResponse{}
	err := utils.UpdateFavoriteNum(req.VideoId, req.ActionType)
	if err != nil {
		resp.Code = config.ErrorCode
		resp.Msg = config.ErrorMsg
		log.Error("UpdateFavoriteNum err:%v", err)
		return &pb.UpdateFavoriteCountRsp{CommonRsp: resp}, err
	}
	resp.Code = config.SuccessCode
	resp.Msg = config.SuccessMsg
	return &pb.UpdateFavoriteCountRsp{CommonRsp: resp}, nil
}
func (VideoService) UpdateCommentCount(ctx context.Context, req *pb.UpdateCommentCountReq) (*pb.UpdateCommentCountRsp, error) {
	resp := &pb.CommonResponse{}
	err := utils.UpdateCommentNum(req.VideoId, req.ActionType)
	if err != nil {
		resp.Code = config.ErrorCode
		resp.Msg = config.ErrorMsg
		log.Error("UpdateCommentNum err:%v", err)
		return &pb.UpdateCommentCountRsp{CommonRsp: resp}, err
	}
	resp.Code = config.SuccessCode
	resp.Msg = config.SuccessMsg
	return &pb.UpdateCommentCountRsp{CommonRsp: resp}, nil
}

func (VideoService) GetVideoInfoList(ctx context.Context, req *pb.GetVideoInfoListReq) (*pb.GetVideoInfoListRsp, error) {
	videosInfo := make([]*pb.VideoInfo, 0)
	for i, v := range req.VideoId {
		videoInfo, err := utils.GetVideoInfo(v)
		if err != nil {
			log.Error("GetVideoInfo err:%v,videoId:%v", err, i)
			continue
		}
		videosInfo = append(videosInfo, videoInfo)
	}
	return &pb.GetVideoInfoListRsp{VideoInfoList: videosInfo}, nil
}

func (v VideoService) GetFavoriteVideoList(ctx context.Context, req *pb.GetFavoriteVideoListReq) (*pb.GetFavoriteVideoListRsp, error) {
	userResp, err := utils.GetFavoriteSvrClient().GetFavoriteList(ctx, &pb.GetFavoriteListRequest{UserId: req.UserId})
	if err != nil {
		log.Errorf("GetFavoriteList err:%v", err)
	}
	resp, err := v.GetVideoInfoList(ctx, &pb.GetVideoInfoListReq{VideoId: userResp.VideoIdList})
	if err != nil {
		log.Errorf("GetVideoInfoList err:%v", err)
	}
	return &pb.GetFavoriteVideoListRsp{VideoList: resp.VideoInfoList}, nil
}

func (VideoService) GetPublishVideoList(ctx context.Context, req *pb.GetPublishVideoListRequest) (*pb.GetPublishVideoListResponse, error) {
	resp, err := utils.GetVideosByUserId(req.UserID)
	if err != nil {
		log.Error("GetVideosByUserId err:%v", err)
	}
	return &pb.GetPublishVideoListResponse{VideoList: resp}, nil
}

func (VideoService) GetFeedList(ctx context.Context, req *pb.GetFeedListRequest) (*pb.GetFeedListResponse, error) {
	videoList, err := utils.GetFeedVideo(req.CurrentTime)
	if err != nil {
		log.Error("GetFeedVideo err:%v", err)
		return &pb.GetFeedListResponse{}, err
	}
	nextTime := time.Now().UnixNano() / int64(time.Millisecond)
	if len(videoList) == 20 {
		nextTime = videoList[len(videoList)-1].PublishTime.UnixNano() / int64(time.Millisecond)
	}
	var VideoInfoList []*pb.VideoInfo
	for _, video := range videoList {
		VideoInfoList = append(VideoInfoList, &pb.VideoInfo{
			Id:            video.Id,
			AuthorId:      video.AuthorId,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false,
			Title:         video.Title,
		})
	}
	resp := &pb.GetFeedListResponse{
		NextTime:  nextTime,
		VideoList: VideoInfoList,
	}

	log.Infof("GetFeedList resp:%+v", resp)
	return resp, nil
}
