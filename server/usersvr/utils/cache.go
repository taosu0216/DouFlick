package utils

import (
	"context"
	"go.uber.org/zap"
	"strconv"
	"usersvr/middleware/cache"
)

func CacheGetUserInfoById(userId int64) (User, error) {
	//hashTable的数据结构,table名是Prefix+userid,然后里面有很多键值对,构成了data这个map[string]string
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)
	data, err := cache.GetRedisCli().HGetAll(context.Background(), userKey).Result()
	if err != nil {
		zap.L().Error("CacheGetUser error", zap.Error(err))
		return User{}, err
	}
	if len(data) == 0 {
		zap.L().Error("CacheGetUser error", zap.Error(err))
		return User{}, err
	}

	user := User{}
	user.Id, _ = strconv.ParseInt(data["id"], 10, 64)
	user.Name = data["user_name"]
	user.Password = data["password"]
	user.Follow, _ = strconv.ParseInt(data["follow_count"], 10, 64)
	user.Follower, _ = strconv.ParseInt(data["follower_count"], 10, 64)
	user.Avatar = data["avatar"]
	user.BackgroundImage = data["background_image"]
	user.Signature = data["signature"]
	user.TotalFav, _ = strconv.ParseInt(data["total_favorited"], 10, 64)
	user.FavCount, _ = strconv.ParseInt(data["favorite_count"], 10, 64)
	return user, nil
}
func CacheSetUser(u User) {
	userKey := userKeyPrefix + strconv.FormatInt(u.Id, 10)
	if err := cache.GetRedisCli().HSet(context.Background(), userKey, map[string]interface{}{
		"id":               u.Id,
		"user_name":        u.Name,
		"password":         u.Password,
		"follow_count":     u.Follow,
		"follower_count":   u.Follower,
		"avatar":           u.Avatar,
		"background_image": u.BackgroundImage,
		"signature":        u.Signature,
		"total_favorited":  u.TotalFav,
		"favorite_count":   u.FavCount,
	}).Err(); err != nil {
		zap.L().Error("CacheSetUser error", zap.Error(err))
		return
	}
}
func CacheInsertUser(u User) {
	userKey := userKeyPrefix + strconv.FormatInt(u.Id, 10)
	if err := cache.GetRedisCli().HSet(context.Background(), userKey, map[string]interface{}{
		"id":               u.Id,
		"user_name":        u.Name,
		"password":         u.Password,
		"follow_count":     u.Follow,
		"follower_count":   u.Follower,
		"avatar":           u.Avatar,
		"background_image": u.BackgroundImage,
		"signature":        u.Signature,
		"total_favorited":  u.TotalFav,
		"favorite_count":   u.FavCount,
	}).Err(); err != nil {
		zap.L().Error("CacheSetUser error", zap.Error(err))
		return
	}
	return
}
func CacheUpdateMyTotalFavouritedCount(userid int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userid, 10)
	if err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "total_favorited", num).Err(); err != nil {
		zap.L().Error("CacheUpdateMyTotalFavouritedCount error", zap.Error(err))
		return err
	}
	return nil
}
func CacheUpdateMyFavouriteVideoCount(userid int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userid, 10)
	if err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "favorite_count", num).Err(); err != nil {
		zap.L().Error("CacheUpdateMyFavouriteVideoCount error", zap.Error(err))
		return err
	}
	return nil
}
func CacheUpdateMyFollowCount(userid int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userid, 10)
	if err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "follow_count", num).Err(); err != nil {
		zap.L().Error("CacheUpdateMyFollowCount error", zap.Error(err))
		return err
	}
	return nil
}
func CacheUpdateMyFollowerCount(userid int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userid, 10)
	if err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "follower_count", num).Err(); err != nil {
		zap.L().Error("CacheUpdateMyFollowerCount error", zap.Error(err))
		return err
	}
	return nil
}
