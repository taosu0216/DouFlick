package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/gin-gonic/gin"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"strconv"
)

type DouFlickRelationListResponse struct {
	StatusCode int32          `json:"status_code"`
	StatusMsg  string         `json:"status_msg,omitempty"`
	UserList   []*pb.UserInfo `json:"user_list,omitempty"`
}

// RelationAction 关注相关的操作
func RelationAction(ctx *gin.Context) {
	tokens, _ := ctx.Get("UserId")
	tokenUserId := tokens.(int64)

	toUserId := ctx.Query("to_user_id")
	toUid, err := strconv.ParseInt(toUserId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	actionStr := ctx.Query("action_type")

	actionType, err := strconv.ParseInt(actionStr, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Infof("RelationAction tokenUserId:%d, toUid:%d, actionType:%d", tokenUserId, toUid, actionType)

	//执行操作(relation)
	_, err = utils.GetRelationSvrClient().RelationAction(ctx, &pb.RelationActionReq{
		ActionType:   actionType,
		SelfUserId:   tokenUserId,
		TargetUserId: toUid,
	})
	if err != nil {
		log.Errorf("RelationAction error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//更新用户自己的关注数(user)
	_, err = utils.GetUserSvrClient().UpdateMyFollowCount(ctx, &pb.UpdateMyFollowCountRequest{
		UserId:     tokenUserId,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("UpdateMyFollowCount error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//更新被关注的人的粉丝数(user)
	_, err = utils.GetUserSvrClient().UpdateMyFollowerCount(ctx, &pb.UpdateMyFollowerCountRequest{
		UserId:     toUid,
		ActionType: actionType,
	})
	if err != nil {
		log.Errorf("UpdateMyFollowerCount error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", nil)
}

// GetFollowList 获取我关注的人的列表
func GetFollowList(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	uid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//获取我关注的人的列表信息 (relation)
	idList, err := utils.GetRelationSvrClient().GetFollowList(ctx, &pb.GetFollowListReq{
		UserId: uid,
	})
	if err != nil {
		log.Errorf("GetFollowList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//根据idList去拿取对应信息 (user)
	userList, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.UserInfoListRequest{IdList: idList.UserList})
	if err != nil {
		log.Errorf("GetUserInfoList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", &DouFlickRelationListResponse{UserList: userList.UserInfoList})
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(ctx *gin.Context) {
	UserId := ctx.Query("user_id")
	uid, err := strconv.ParseInt(UserId, 10, 64)
	if err != nil {
		log.Errorf("GetFollowerList ParseInt error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	//获取粉丝列表的id (relation)
	idList, err := utils.GetRelationSvrClient().GetFansList(ctx, &pb.GetFansListReq{
		UserId: uid,
	})
	if err != nil {
		log.Errorf("GetFansList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	//根据id拿info(user)
	infos, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.UserInfoListRequest{IdList: idList.FollowerList})
	if err != nil {
		log.Errorf("GetUserInfoList error : %s", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", &DouFlickRelationListResponse{UserList: infos.UserInfoList})
}
