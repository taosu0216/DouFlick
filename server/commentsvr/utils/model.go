package utils

import "time"

var (
	CommentInfoPrefix = "DouFlick:comment:"
	VideoInfoPrefix   = "DouFlick:video:"
)

type Comment struct {
	ID          int64     `gorm:"column:id;primary_key;" json:"id"`
	UserId      int64     `gorm:"column:user_id" json:"user_id"`
	VideoId     int64     `gorm:"video_id" json:"video_id"`
	CommentText string    `gorm:"comment_text" json:"comment_text"`
	CreateTime  time.Time `gorm:"create_time" json:"create_time"`
}

func (c *Comment) TableName() string {
	return "t_comment"
}
