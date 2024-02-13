package service

import (
	"commentsvr/config"
	"commentsvr/log"
	"commentsvr/utils"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"golang.org/x/net/context"
)

type CommentService struct {
	pb.UnimplementedCommentServiceServer
}

func (CommentService) GetCommentList(ctx context.Context, req *pb.CommentListRequest) (*pb.CommentListResponse, error) {
	comments, err := utils.GetComments(req.VideoId)
	if err != nil {
		log.Error("GetComments error", err)
		return nil, err
	}

	userIds := make([]int64, len(comments))
	for i, comment := range comments {
		userIds[i] = comment.UserId
	}

	userInfoLists, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.UserInfoListRequest{
		IdList: userIds,
	})
	if err != nil {
		log.Error("GetUserInfoList error", err)
		return nil, err
	}

	userMap := make(map[int64]*pb.UserInfo)
	for _, userInfo := range userInfoLists.UserInfoList {
		userMap[userInfo.Id] = userInfo
	}

	commentListResp := &pb.CommentListResponse{
		CommentList: make([]*pb.Comment, len(comments)),
	}
	for i, userinfo := range comments {
		v := &pb.Comment{
			Id:          userinfo.ID,
			User:        userMap[userinfo.UserId],
			CommentText: userinfo.CommentText,
			CreateTime:  userinfo.CreateTime.Format(config.DefaultTime),
		}
		commentListResp.CommentList[i] = v
	}
	return commentListResp, nil
}

func (CommentService) CommentAction(ctx context.Context, req *pb.CommentRequest) (*pb.CommentResponse, error) {
	if req.ActionType == 1 {
		//增加评论
		comment, err := utils.CommentAdd(req.UserId, req.VideoId, req.CommentText)
		if err != nil {
			log.Error("CommentAdd error", err)
			return nil, err
		}
		getUserInfoRsp, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.UserInfoRequest{
			Id: req.UserId,
		})
		if err != nil {
			log.Error("GetUserInfo error", err)
			return nil, err
		}
		result := &pb.CommentResponse{Comment: &pb.Comment{
			Id:          comment.ID,
			User:        getUserInfoRsp.UserInfo,
			CommentText: comment.CommentText,
			CreateTime:  comment.CreateTime.Format(config.DefaultTime),
		}}
		return result, nil
	} else {
		// TODO: 这里都需要重新看
		err := utils.CommentDelete(req.VideoId, req.CommentId)
		if err != nil {
			log.Error("CommentDelete error", err)
			return nil, err
		}
		return &pb.CommentResponse{
			Comment: nil,
		}, nil
	}
}
