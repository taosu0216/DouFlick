package utils

import (
	"fmt"
	"relationsvr/init/db"

	"gorm.io/gorm"

	"go.uber.org/zap"
)

func DBGetFollowList(userId int64) (list []int64, err error) {
	mysqlDB := db.GetMySqlDB()
	relationList := make([]*Relation, 0)
	err = mysqlDB.Where("follower_id = ?", userId).Find(&relationList).Error
	if err != nil {
		zap.L().Error("DBGetFollowList error", zap.Error(err))
		return nil, err
	}
	for _, relation := range relationList {
		list = append(list, relation.Follow)
	}
	return list, nil
}

func DBGetFansList(userId int64) (list []int64, err error) {
	mysqlDB := db.GetMySqlDB()
	relationList := make([]*Relation, 0)
	err = mysqlDB.Where("follow_id = ?", userId).Find(&relationList).Error
	if err != nil {
		zap.L().Error("DBGetFansList error", zap.Error(err))
		return nil, err
	}
	for _, relation := range relationList {
		list = append(list, relation.Follower)
	}
	return list, nil
}

func DBIsFollow(userId, targetId int64) (bool, error) {
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Where("follow_id = ? and follower_id = ?", targetId, userId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return false, nil
		}
		zap.L().Error("DBIsFollow error", zap.Error(err))
		return false, err
	}
	return true, nil
}

func DBFollowAction(userId, targetId int64) error {
	mysqlDB := db.GetMySqlDB()
	relation := &Relation{
		Follow:   targetId,
		Follower: userId,
	}
	err := mysqlDB.Where("follow_id = ? and follower_id = ?", targetId, userId).First(&Relation{}).Error
	if err == nil {
		return fmt.Errorf("you have followed this user")
	}
	if err != nil && err.Error() != gorm.ErrRecordNotFound.Error() {
		zap.L().Error("DBFollowAction error", zap.Error(err))
		return err
	}
	err = mysqlDB.Create(relation).Error
	if err != nil {
		zap.L().Error("DBFollowAction error", zap.Error(err))
		return err
	}
	return nil
}

func DBUnFollowAction(userId, targetId int64) error {
	mysqlDB := db.GetMySqlDB()
	err := mysqlDB.Where("follow_id = ? and follower_id = ?", targetId, userId).Delete(&Relation{}).Error
	if err != nil {
		zap.L().Error("DBUnFollowAction error", zap.Error(err))
		return err
	}
	return nil
}
