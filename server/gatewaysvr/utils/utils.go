package utils

import "github.com/taosu0216/DouFlick/pkg/pb"

var (
	UserSvrClient    pb.UserServiceClient
	CommentSvrClient pb.CommentServiceClient
)

func GetUserSvrClient() pb.UserServiceClient {
	return UserSvrClient
}

func GetCommentSvrClient() pb.CommentServiceClient {
	return CommentSvrClient
}
