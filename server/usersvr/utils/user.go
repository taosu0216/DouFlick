package utils

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GetUserInfoByAny(u interface{}) (user User, err error) {
	switch u := u.(type) {
	case int64:
		user, err = CacheGetUserInfoById(u)
		if err == nil {
			return user, nil
		}
		user, err = DbGetUserByUserId(u)
	case string:
		user, err = DbGetUserByUserName(u)
	default:
		err = errors.New("nil error")
	}
	return user, err
}

func GetUserList(useridList []int64) ([]*User, error) {
	users, err := DbGetUserList(useridList)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func IsUserNameExist(name string) (bool, error) {
	user, err := DbGetUserByUserName(name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("##################", err)
		zap.L().Error("IsUserNameExist error", zap.Error(err))
		return false, err
	}
	if user.Id != 0 {
		return true, nil
	}
	return false, nil
}

func CreateUserToMysqlAndCache(username, password string) (*User, error) {
	//先存入数据库
	user, err := DbInsertUser(username, password)
	if err != nil {
		//上一层已经封装过error,所以直接返回
		return nil, err
	}
	//再存入redis
	go CacheInsertUser(user)
	return &user, nil
}

func UpdateMyTotalFavouritedCount(userId int64, updateType int64) error {
	var num int64
	if updateType == 1 {
		num = 1
	} else {
		num = -1
	}

	//更新mysql
	err := DbUpdateMyTotalFavouritedCount(userId, num)
	if err != nil {
		return err
	}

	//更新redis
	err = CacheUpdateMyTotalFavouritedCount(userId, num)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMyFavouriteVideoCount(userId, updateNum int64) error {
	var num int64
	if updateNum == 1 {
		num = 1
	} else {
		num = -1
	}

	//更新mysql
	err := DbUpdateMyFavouriteVideoCount(userId, num)
	if err != nil {
		return err
	}

	//更新redis
	err = CacheUpdateMyFavouriteVideoCount(userId, num)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMyFollowCount(userId, updateNum int64) error {
	var num int64
	if updateNum == 1 {
		num = 1
	} else {
		num = -1
	}

	//更新mysql
	err := DbUpdateMyFollowCount(userId, num)
	if err != nil {
		return err
	}

	//更新redis
	err = CacheUpdateMyFollowCount(userId, num)
	if err != nil {
		return err
	}
	return nil
}

func UpdateMyFollowerCount(userId, updateNum int64) error {
	var num int64
	if updateNum == 1 {
		num = 1
	} else {
		num = -1
	}

	//更新mysql
	err := DbUpdateMyFollowerCount(userId, num)
	if err != nil {
		return err
	}

	//更新redis
	err = CacheUpdateMyFollowerCount(userId, num)
	if err != nil {
		return err
	}
	return nil
}
