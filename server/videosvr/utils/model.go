package utils

import "time"

type Video struct {
	Id            int64     `gorm:"column:id; primary_key;"` // video_id
	AuthorId      int64     `gorm:"column:author_id;"`       // 谁发布的
	PlayUrl       string    `gorm:"column:play_url;"`        // videoURL
	CoverUrl      string    `gorm:"column:cover_url;"`       // picURL
	FavoriteCount int64     `gorm:"column:favorite_count;"`  // 点赞数
	CommentCount  int64     `gorm:"column:comment_count;"`   // 评论数
	PublishTime   time.Time `gorm:"column:publish_time;"`    // 发布时间
	Title         string    `gorm:"column:title;"`           // 标题
}

func (Video) TableName() string {
	return "t_video"
}
