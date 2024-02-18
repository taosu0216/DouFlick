package controller

import (
	"fmt"
	"gatewaysvr/config"
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/gin-gonic/gin"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"path/filepath"
	"strconv"
)

type DouFlickVideoActionResp struct {
	StatusCode int32  `json:"status_code" protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3"`
	StatusMsg  string `json:"status_msg,omitempty" protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3"`
}

type DouFlickVideoListResp struct {
	StatusCode int32    `json:"status_code" protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3"`
	StatusMsg  string   `json:"status_msg,omitempty" protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

type Video struct {
	Id            int64        `json:"id"`
	Author        *pb.UserInfo `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount int64        `json:"favorite_count"`
	CommentCount  int64        `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

func Publish(c *gin.Context) {
	userId, _ := c.Get("UserId")
	title := c.PostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}
	fmt.Println("data", data)
	filename := filepath.Base(data.Filename)
	random := utils.RandomStr()
	filename = fmt.Sprintf("%s_%s", random, filename)
	videoPath := config.GetGlobalConfig().VideoPath
	saveFile := filepath.Join(videoPath, filename)
	log.Infof("videoPath:%v", videoPath)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		response.Fail(c, err.Error(), nil)
		return
	}

	//调用video模块的rpc方法
	fmt.Println(saveFile)
	_, err = utils.GetVideoSvrClient().PublishVideo(c, &pb.PublishVideoRequest{
		UserId:   userId.(int64),
		Title:    title,
		SaveFile: saveFile,
	})
	if err != nil {
		log.Errorf("utils.GetVideoSvrClient().PublishVideo err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	response.Success(c, "success", &DouFlickVideoActionResp{})

}

// GetVideoList 获取自己发布视频列表
func GetVideoList(ctx *gin.Context) {
	tokenUid, _ := ctx.Get("UserId")
	id := ctx.Query("user_id")
	userId := tokenUid.(int64)
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	getPublishVideoList, err := utils.GetVideoSvrClient().GetPublishVideoList(ctx, &pb.GetPublishVideoListRequest{TokenUserId: userId, UserID: uid})
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//获取用户自己的信息
	UserInfo, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.UserInfoRequest{Id: tokenUid.(int64)})
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}
	video := make([]*Video, 0, len(getPublishVideoList.VideoList))
	for _, v := range getPublishVideoList.VideoList {
		video = append(video, &Video{
			Id:            v.Id,
			Author:        UserInfo.UserInfo,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			// TODO:
			IsFavorite: false, // 自己是否喜欢
		})
	}
	response.Success(ctx, "success", &DouFlickVideoListResp{
		VideoList: video,
	})
}
