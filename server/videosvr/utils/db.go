package utils

import (
	"github.com/taosu0216/DouFlick/pkg/pb"
	"gorm.io/gorm"
	"time"
	"videosvr/init/db"
)

func DBInsertVideo(userId int64, videoUrl string, picUrl string, title string) error {
	mysqlDB := db.GetMySqlDB()
	video := &Video{AuthorId: userId, PlayUrl: videoUrl, CoverUrl: picUrl, FavoriteCount: 0, CommentCount: 0, PublishTime: time.Now(), Title: title}
	err := mysqlDB.Model(&Video{}).Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}

func DBUpdateFavoriteNum(videoId int64, actionType int64) error {
	mysqlDB := db.GetMySqlDB()
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}
	err := mysqlDB.Model(&Video{}).Where("id = ?", videoId).Update("favorite_count", gorm.Expr("favorite_count + ?", num)).Error
	if err != nil {
		return err
	}
	return nil
}

func DBUpdateCommentNum(videoId int64, actionType int64) error {
	mysqlDB := db.GetMySqlDB()
	var num int64
	if actionType == 1 {
		num = 1
	} else {
		num = -1
	}
	err := mysqlDB.Model(&Video{}).Where("id = ?", videoId).Update("comment_count", gorm.Expr("comment_count + ?", num)).Error
	if err != nil {
		return err
	}
	return nil
}

func DBGetVideoInfo(videoId int64) (*pb.VideoInfo, error) {
	mysqlDB := db.GetMySqlDB()
	var video Video
	err := mysqlDB.Where("id = ?", videoId).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &pb.VideoInfo{Id: video.Id, AuthorId: video.AuthorId, PlayUrl: video.PlayUrl, CoverUrl: video.CoverUrl, FavoriteCount: video.FavoriteCount, CommentCount: video.CommentCount, IsFavorite: false, Title: video.Title}, nil
}

func DBGetVideosByUserId(userId int64) ([]Video, error) {
	mysqlDB := db.GetMySqlDB()
	var videos []Video
	err := mysqlDB.Model(&Video{}).Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func DBGetFeedVideo(curTime int64) ([]Video, error) {
	mysqlDB := db.GetMySqlDB()
	var videos []Video
	err := mysqlDB.Model(&Video{}).Where("publish_time < ?", curTime).Limit(20).Order("id DESC").Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
