package controller

import (
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"go.uber.org/zap"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/taosu0216/DouFlick/pkg/pb"
)

type DouFlickCommentListResp struct {
	Status      int32         `json:"status"`
	StatusMsg   string        `json:"status_msg,omitempty"`
	CommentList []*pb.Comment `json:"comment_list,omitempty"`
}

type DouFlickAddCommentResp struct {
	Status    int32       `json:"status"`
	StatusMsg string      `json:"status_msg,omitempty"`
	Comment   *pb.Comment `json:"comment"`
}

func GetCommentList(c *gin.Context) {
	videoId, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		zap.L().Error("videoId error", zap.Error(err))
		response.Fail(c, "video_id invalid", nil)
		return
	}
	request := &pb.CommentListRequest{VideoId: videoId}
	resp, err := utils.GetCommentSvrClient().GetCommentList(c, request)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	response.Success(c, "success", &DouFlickCommentListResp{
		CommentList: resp.CommentList,
	})
}

func AddComment(c *gin.Context) {

	tokenUids, _ := c.Get("UserId")

	tokenUid := tokenUids.(int64)

	videoId, err := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
	comment := c.PostForm("comment_text")
	actionTypeStr := c.PostForm("action_type")
	commentID := c.PostForm("comment_id")

	commentId, err := strconv.ParseInt(commentID, 10, 64)
	if err != nil {
		response.Fail(c, "comment_id invalid", nil)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		response.Fail(c, "action_type invalid", nil)
		return
	}
	resq := &pb.CommentRequest{
		VideoId:     videoId,
		UserId:      tokenUid,
		CommentText: comment,
		ActionType:  actionType,
		CommentId:   commentId,
	}

	//发布评论
	resp, err := utils.GetCommentSvrClient().CommentAction(c, resq)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	//视频评论+1
	_, err = utils.GetVideoSvrClient().UpdateCommentCount(c, &pb.UpdateCommentCountReq{VideoId: videoId, ActionType: actionType})
	// TODO: video模块未完成
	userinfo, err := utils.GetUserSvrClient().GetUserInfo(c, &pb.UserInfoRequest{Id: tokenUid})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	resp.Comment.User = userinfo.UserInfo
	response.Success(c, "success", &DouFlickAddCommentResp{Comment: resp.Comment})
}
