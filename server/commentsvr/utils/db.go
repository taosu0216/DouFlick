package utils

import (
	"commentsvr/init/db"
	"go.uber.org/zap"
	"time"
)

func DbCommentAdd(userId, videoId int64, comment_text string) (*Comment, error) {
	sqlDB := db.GetMySqlDB()

	nowTime := time.Now()
	comment := &Comment{
		UserId:      userId,
		VideoId:     videoId,
		CommentText: comment_text,
		CreateTime:  nowTime,
	}
	result := sqlDB.Create(comment)

	if result.Error != nil {
		zap.L().Error("DbCommentAdd failed", zap.Error(result.Error))
		return nil, result.Error
	}
	return comment, nil
}

func DbCommentDelete(commentId int64) error {
	sqlDB := db.GetMySqlDB()

	err := sqlDB.Model(&Comment{}).Where("id = ?", commentId).Delete(&Comment{}).Error
	if err != nil {
		zap.L().Error("DbCommentDelete failed", zap.Error(err))
		return err
	}
	return nil
}

func DbGetComments(videoId int64) ([]*Comment, error) {
	sqlDB := db.GetMySqlDB()

	var comments []*Comment
	err := sqlDB.Where("video_id = ?", videoId).Order("id DESC").Find(&comments).Error
	if err != nil {
		zap.L().Error("DbGetComments failed", zap.Error(err))
		return nil, err
	}
	return comments, nil
}
