package utils

import (
	"commentsvr/init/cache"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"strconv"
)

// CacheSetComment 给某一个video添加评论，评论以hash 形式存储
func CacheSetComment(comment *Comment) error {
	videoKey := VideoInfoPrefix + strconv.FormatInt(comment.VideoId, 10)
	redisCli := cache.GetRedisCli()

	// 设置过期时间
	expired := cache.ValueExpire

	commentBytes, err := json.Marshal(comment)
	if err != nil {
		zap.L().Error("json marshal comment failed", zap.Error(err))
		return err
	}
	commentIdStr := strconv.FormatInt(comment.ID, 10)
	err = redisCli.HSet(context.Background(), videoKey, commentIdStr, string(commentBytes)).Err()
	if err != nil {
		zap.L().Error("redis hset failed", zap.Error(err))
		return err
	}
	redisCli.Expire(context.Background(), videoKey, expired)
	return nil
}

func CacheDeleteComment(keyList []string, VideoID int64) error {
	videoKey := VideoInfoPrefix + strconv.FormatInt(VideoID, 10)
	redisCli := cache.GetRedisCli()

	if err := redisCli.HDel(context.Background(), videoKey, keyList...).Err(); err != nil {
		zap.L().Error("redis hdel failed", zap.Error(err))
		return err
	}
	return nil
}

func CacheGetComments(videoId int64) ([]*Comment, error) {
	videoKey := VideoInfoPrefix + strconv.FormatInt(videoId, 10)
	redisCli := cache.GetRedisCli()

	data, err := redisCli.HGetAll(context.Background(), videoKey).Result()
	if err != nil {
		zap.L().Error("CacheGetComments error", zap.Error(err))
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}

	comments := make([]*Comment, 0, len(data))
	for _, v := range data {
		comment := &Comment{}
		if err := json.Unmarshal([]byte(v), comment); err != nil {
			zap.L().Error("json unmarshal comment failed", zap.Error(err))
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
