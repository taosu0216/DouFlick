package service

import (
	"fmt"
	"relationsvr/config"
	"relationsvr/log"
	"relationsvr/utils"
	"strconv"

	"github.com/taosu0216/DouFlick/pkg/pb"
	"golang.org/x/net/context"
)

/*
  //执行关注或取消关注操作
  rpc RelationAction(RelationActionReq) returns (RelationActionResp);

  //获取用户关注列表
  rpc GetFollowList(GetFollowListReq) returns (GetFollowListResp);

  //获取用户粉丝列表
  rpc GetFansList(GetFansListReq) returns (GetFansListResp);

  //检查用户是否关注了特定用户
  rpc IsFollowDict(IsFollowDictReq) returns (IsFollowDictResp);
*/

type RelationService struct {
	pb.UnimplementedRelationServiceServer
}

func (RelationService) RelationAction(ctx context.Context, req *pb.RelationActionReq) (*pb.RelationActionResp, error) {
	//TODO:在网关层就进行是否关注自己的判断
	if req.SelfUserId == req.TargetUserId {
		return &pb.RelationActionResp{
			CommonResponse: &pb.CommonResponse{
				Code: config.ErrorCode,
				Msg:  "you can't follow yourself",
			},
		}, fmt.Errorf("you can't follow yourself")
	}
	if req.ActionType == 1 {
		err := utils.FollowAction(req.SelfUserId, req.TargetUserId)
		if err != nil {
			log.Error("RelationAction error", err)
			return nil, err
		}
	} else {
		err := utils.UnFollowAction(req.SelfUserId, req.TargetUserId)
		if err != nil {
			log.Error("RelationAction error", err)
			return nil, err
		}
	}
	return &pb.RelationActionResp{
		CommonResponse: &pb.CommonResponse{
			Code: config.SuccessCode,
			Msg:  config.SuccessMsg,
		},
	}, nil
}

func (RelationService) GetFollowList(ctx context.Context, req *pb.GetFollowListReq) (*pb.GetFollowListResp, error) {
	myFollowList, err := utils.GetFollowList(req.UserId)
	if err != nil {
		log.Error("GetFollowList error", err)
		return nil, err
	}
	return &pb.GetFollowListResp{UserList: myFollowList}, nil
}

func (RelationService) GetFansList(ctx context.Context, req *pb.GetFansListReq) (*pb.GetFansListResp, error) {
	fansList, err := utils.GetFansList(req.UserId)
	if err != nil {
		log.Error("GetFansList error", err)
		return nil, err
	}
	return &pb.GetFansListResp{FollowerList: fansList}, nil
}

func (RelationService) IsFollow(ctx context.Context, req *pb.IsFollowDictReq) (*pb.IsFollowDictResp, error) {
	dict := make(map[string]bool)
	for _, unit := range req.FollowUnit {
		isFollow, err := utils.IsFollow(unit.SelfUserId, unit.TargetUserId)
		if err != nil {
			log.Error("IsFollowDict error", err)
			return nil, err
		}
		isFollowKey := strconv.FormatInt(unit.SelfUserId, 10) + "_" + strconv.FormatInt(unit.TargetUserId, 10)
		dict[isFollowKey] = isFollow
	}
	return &pb.IsFollowDictResp{IsFollow: dict}, nil
}
