package utils

import (
	"github.com/taosu0216/DouFlick/pkg/pb"
	"videosvr/log"
)

func InsertVideo(userId int64, videoUrl string, picUrl string, title string) error {
	return DBInsertVideo(userId, videoUrl, picUrl, title)
}
func UpdateFavoriteNum(videoId int64, actionType int64) error {
	return DBUpdateFavoriteNum(videoId, actionType)
}
func UpdateCommentNum(videoId int64, actionType int64) error {
	return DBUpdateCommentNum(videoId, actionType)
}
func GetVideoInfo(videoId int64) (*pb.VideoInfo, error) {
	info, err := DBGetVideoInfo(videoId)
	if err != nil {
		return nil, err
	}
	return info, nil
}
func GetVideosByUserId(userId int64) ([]*pb.VideoInfo, error) {
	infos, err := DBGetVideosByUserId(userId)
	if err != nil {
		return nil, err
	}
	resp := make([]*pb.VideoInfo, 0)
	for _, video := range infos {
		resp = append(resp, &pb.VideoInfo{
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
	return resp, nil
}

func GetFeedVideo(curTime int64) ([]Video, error) {
	videos, err := DBGetFeedVideo(curTime)
	if err != nil {
		return nil, err
	}
	log.Info("GetVideoListByFeed", videos)
	return videos, nil
}
