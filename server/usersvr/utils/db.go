package utils

import (
	"errors"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"usersvr/middleware/db"
)

func DbGetUserList(useridlist []int64) ([]*User, error) {
	mysqldb := db.GetMySqlDB()
	var users []*User
	err := mysqldb.Table("t_user").Where("id in ?", useridlist).Find(&users).Error
	if err != nil {
		zap.L().Error("DbGeuUserList error", zap.Error(err))
		return nil, err
	}
	return users, nil
}
func DbGetUserByUserId(id int64) (User, error) {
	sqlDB := db.GetMySqlDB()
	user := User{}
	err := sqlDB.Where("id = ?", id).First(&user).Error
	if err != nil {
		zap.L().Error("DbGetUserByUserId error", zap.Error(err))
		return User{}, err
	}
	return user, nil
}
func DbGetUserByUserName(name string) (User, error) {
	sqlDB := db.GetMySqlDB()
	user := User{}
	err := sqlDB.Where("user_name = ?", name).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		zap.L().Error("DbGetUserByUserName error", zap.Error(err))
		return User{}, err
	}
	return user, err
}
func DbInsertUser(username, password string) (User, error) {
	sqlDB := db.GetMySqlDB()
	hashPasswd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := User{
		Name:            username,
		Password:        string(hashPasswd),
		Follow:          0,
		Follower:        0,
		TotalFav:        0,
		FavCount:        0,
		Avatar:          "https://cdn.jsdelivr.net/gh/taosu0216/picgo/20230821231539.png",
		BackgroundImage: "https://cdn.jsdelivr.net/gh/taosu0216/picgo/20231022175028.png",
		Signature:       "要开心哦~",
	}
	//这里的迁移是自动迁移,所以不需要手动创建表
	//自动迁移指定表的结构是  .Model(&User{})这里的User{}是结构体,不是表名
	result := sqlDB.Model(&User{}).Create(&user)
	if result.Error != nil {
		zap.L().Error("DbInsertUser error", zap.Error(result.Error))
		return User{}, result.Error
	}
	return user, nil
}
func DbUpdateMyTotalFavouritedCount(userId int64, num int64) error {
	sqlDB := db.GetMySqlDB()
	//根据User{}查找对应的表名,然后更新total_favorited字段
	err := sqlDB.Model(&User{}).Where("id = ?", userId).Update("total_favorited", gorm.Expr("total_favorited + ?", num)).Error
	if err != nil {
		zap.L().Error("DbUpdateMyTotalFavouritedCount error", zap.Error(err))
		return err
	}
	return nil
}
func DbUpdateMyFavouriteVideoCount(userId, num int64) error {
	sqlDB := db.GetMySqlDB()
	err := sqlDB.Model(&User{}).Where("id = ?", userId).Update("favorite_count", gorm.Expr("favorite_count + ?", num)).Error
	if err != nil {
		zap.L().Error("DbUpdateMyFavouriteVideoCount error", zap.Error(err))
		return err
	}
	return nil
}
func DbUpdateMyFollowCount(userId, num int64) error {
	sqlDB := db.GetMySqlDB()
	err := sqlDB.Model(&User{}).Where("id = ?", userId).Update("follow_count", gorm.Expr("follow_count + ?", num)).Error
	if err != nil {
		zap.L().Error("DbUpdateMyFollowCount error", zap.Error(err))
		return err
	}
	return nil
}
func DbUpdateMyFollowerCount(userId, num int64) error {
	sqlDB := db.GetMySqlDB()
	err := sqlDB.Model(&User{}).Where("id = ?", userId).Update("follower_count", gorm.Expr("follower_count + ?", num)).Error
	if err != nil {
		zap.L().Error("DbUpdateMyFollowerCount error", zap.Error(err))
		return err
	}
	return nil
}
