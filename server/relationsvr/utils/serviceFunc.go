package utils

import "go.uber.org/zap"

func GetFollowList(userId int64) ([]int64, error) {
	list, err := DBGetFollowList(userId)
	if err != nil {
		zap.L().Error("DBGetFollowList error", zap.Error(err))
		return nil, err
	}
	return list, err
}

func GetFansList(userId int64) ([]int64, error) {
	list, err := DBGetFansList(userId)
	if err != nil {
		zap.L().Error("DBGetFansList error", zap.Error(err))
		return nil, err
	}
	return list, err
}

func IsFollow(userId, targetId int64) (bool, error) {
	//这个判断是不是自己关注自己,应该在网关层就判断了
	//if userId == targetId {
	//	return true, nil
	//}
	return DBIsFollow(userId, targetId)
}

func FollowAction(userId, targetId int64) error {
	return DBFollowAction(userId, targetId)
}

func UnFollowAction(userId, targetId int64) error {
	return DBUnFollowAction(userId, targetId)
}
