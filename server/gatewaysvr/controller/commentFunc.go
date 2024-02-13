package controller

import (
	"gatewaysvr/response"
	"gatewaysvr/utils"
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
	video_id, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	//TODO: 鉴权
	if err != nil {
		response.Fail(c, "video_id invalid", nil)
		return
	}
	request := &pb.CommentListRequest{VideoId: video_id}
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
	video_id, err := strconv.ParseInt(c.PostForm("video_id"), 10, 64)
	user_id, err := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
	comment := c.PostForm("comment_text")
	actionTypeStr := c.PostForm("action_type")
	comment_id := c.PostForm("comment_id")
	commentId := int64(0)
	if actionTypeStr == "2" {
		commentId, err = strconv.ParseInt(comment_id, 10, 64)
		if err != nil {
			response.Fail(c, "comment_id invalid", nil)
		}
	}
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 32)
	if err != nil {
		response.Fail(c, "action_type invalid", nil)
		return
	}
	resq := &pb.CommentRequest{
		VideoId:     video_id,
		UserId:      user_id,
		CommentText: comment,
		ActionType:  actionType,
		CommentId:   commentId,
	}
	resp, err := utils.GetCommentSvrClient().CommentAction(c, resq)
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	// TODO: video模块未完成
	userinfo, err := utils.GetUserSvrClient().GetUserInfo(c, &pb.UserInfoRequest{Id: user_id})
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	resp.Comment.User = userinfo.UserInfo
	response.Success(c, "success", &DouFlickAddCommentResp{Comment: resp.Comment})
}
