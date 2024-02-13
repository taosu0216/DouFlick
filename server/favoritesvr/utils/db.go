package utils

import (
	"errors"
	"favoritesvr/init/db"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

func DBGetFavoriteList(userId int64) ([]int64, error) {
	var list []*Favorite
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Model(&Favorite{}).Where("user_id = ?", userId).Find(&list).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Error("DBGetFavoriteList error", zap.Error(err))
		return nil, err
	}
	videoIDs := make([]int64, len(list))
	for i, favorite := range list {
		videoIDs[i] = favorite.VideoId
	}
	return videoIDs, nil
}

func DBIsFavoriteDict(userId, videoId int64) (bool, error) {
	var favorite Favorite
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).First(&favorite).Error
	if err != nil || errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Error("DBIsFavoriteDict error", zap.Error(err))
		return false, err
	}
	return true, nil
}

func DBFavoriteAdd(userId, videoId int64) error {
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Create(&Favorite{UserId: userId, VideoId: videoId}).Error
	if err != nil {
		zap.L().Error("DBFavoriteAdd error", zap.Error(err))
		return err
	}
	return nil
}

func DBFavoriteDelete(userId, videoId int64) error {
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Model(&Favorite{}).Where("user_id = ? AND video_id = ?", userId, videoId).Delete(&Favorite{}).Error
	if err != nil {
		zap.L().Error("DBFavoriteDelete error", zap.Error(err))
		return err
	}
	return nil
}
