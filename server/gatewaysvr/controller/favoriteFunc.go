package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/gin-gonic/gin"
	"github.com/taosu0216/DouFlick/pkg/pb"
)

type FavActionParams struct {
	// 暂时没 user_id ，因为客户端出于安全考虑没给出
	Token      string `form:"token" binding:"required"`
	VideoId    int64  `form:"video_id" binding:"required"`
	ActionType int64  `form:"action_type" binding:"required,oneof=1 2"`
}

type FavListParams struct {
	Token  string `form:"token" binding:"required"`
	UserId int64  `form:"user_id" binding:"required"`
}

type DouFlickFavoriteListResponse struct {
	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`
}

// FavoriteAction 执行点赞或取消点赞的操作
func FavoriteAction(c *gin.Context) {
	var favInfo FavActionParams
	err := c.ShouldBindHeader(&favInfo)
	if err != nil {
		log.Errorf("ShouldBindQuery failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}
	tokenUidStr, _ := c.Get("UserId")
	tokenUid := tokenUidStr.(int64)

	//rpc调用favorite模块的点赞方法(favorite)
	_, err = utils.GetFavoriteSvrClient().FavoriteAction(c, &pb.FavoriteActionRequest{
		UserId:     tokenUid,
		VideoId:    favInfo.VideoId,
		ActionType: favInfo.ActionType,
	})
	if err != nil {
		log.Errorf("FavoriteAction failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	//更新video表的点赞数(video)
	_, err = utils.GetVideoSvrClient().UpdateFavoriteCount(c, &pb.UpdateFavoriteCountReq{
		VideoId:    favInfo.VideoId,
		ActionType: favInfo.ActionType,
	})
	if err != nil {
		log.Errorf("UpdateFavoriteCount failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	userInfoResp, err := utils.GetVideoSvrClient().GetVideoInfoList(c, &pb.GetVideoInfoListReq{
		VideoId: []int64{favInfo.VideoId},
	})
	if err != nil {
		log.Errorf("GetVideoInfoList failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}
	authorId := userInfoResp.VideoInfoList[0].AuthorId

	//更新视频作者的获赞数(user)
	_, err = utils.GetUserSvrClient().UpdateMyFavouritedCount(c, &pb.UpdateMyFavouritedCountRequest{
		UserId:     authorId,
		ActionType: favInfo.ActionType,
	})
	if err != nil {
		log.Errorf("UpdateMyFavouritedCount failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	//更新点赞者的点赞视频(user)
	_, err = utils.GetUserSvrClient().UpdateMyFavouriteCount(c, &pb.UpdateMyFavouriteCountRequest{
		UserId:     tokenUid,
		ActionType: favInfo.ActionType,
	})
	if err != nil {
		log.Errorf("UpdateMyFavouriteCount failed, err:%v", err)
		response.Fail(c, err.Error(), nil)
		return
	}

	response.Success(c, "success", nil)
}

func GetFavoriteList(ctx *gin.Context) {
	tokenUidStr, _ := ctx.Get("UserId")
	tokenUid := tokenUidStr.(int64)

	//根据id拿点赞的视频(video)
	favList, err := utils.GetVideoSvrClient().GetFavoriteVideoList(ctx, &pb.GetFavoriteVideoListReq{
		UserId: tokenUid,
	})
	if err != nil {
		log.Errorf("GetFavoriteVideoList failed, err:%v", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	var userIds []int64
	for _, v := range favList.VideoList {
		userIds = append(userIds, v.AuthorId)
	}

	//拿每个视频对应的user信息(user)
	userInfos, err := utils.GetUserSvrClient().GetUserInfoList(ctx, &pb.UserInfoListRequest{
		IdList: userIds,
	})
	if err != nil {
		log.Errorf("GetUserInfoList failed, err:%v", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}

	//把拿到的userinfos打包封装到resp中
	userMap := make(map[int64]*pb.UserInfo)
	for _, v := range userInfos.UserInfoList {
		userMap[v.Id] = v
	}
	videoList := make([]*Video, 0)

	//TODO: 这里逻辑有点乱
	for _, v := range favList.VideoList {
		videoList = append(videoList, &Video{
			Id:            v.Id,
			Title:         v.Title,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Author: &pb.UserInfo{
				Id:              v.AuthorId,
				Name:            userMap[v.AuthorId].Name,
				Avatar:          userMap[v.AuthorId].Avatar,
				FollowCount:     userMap[v.AuthorId].FollowCount,
				FollowerCount:   userMap[v.AuthorId].FollowerCount,
				IsFollow:        userMap[v.AuthorId].IsFollow,
				Background:      userMap[v.AuthorId].Background,
				Signature:       userMap[v.AuthorId].Signature,
				TotalFavourited: userMap[v.AuthorId].TotalFavourited,
				FavouriteCount:  userMap[v.AuthorId].FavouriteCount,
			},
		})
	}
	if err != nil {
		response.Fail(ctx, err.Error(), nil)
		return
	}

	response.Success(ctx, "success", &DouFlickFavoriteListResponse{
		VideoList: videoList,
	})
}
