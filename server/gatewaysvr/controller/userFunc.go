package controller

import (
	"gatewaysvr/log"
	"gatewaysvr/response"
	"gatewaysvr/utils"
	"github.com/gin-gonic/gin"
	"github.com/taosu0216/DouFlick/pkg/pb"
	"go.uber.org/zap"
	"strconv"
)

func UserLogin(ctx *gin.Context) {
	userName := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(userName) > 32 || len(password) > 32 {
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	resp, err := utils.GetUserSvrClient().CheckPassWord(ctx, &pb.CheckPassWordRequest{
		Username: userName,
		Password: password,
	})
	if err != nil {
		zap.L().Error("login error", zap.Error(err))
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", resp)
}
func GetUserInfo(ctx *gin.Context) {
	userid := ctx.Query("user_id")
	uids, _ := ctx.Get("UserId")
	uid := uids.(int64)
	if strconv.FormatInt(uid, 10) != userid {
		response.Fail(ctx, "token error", nil)
		return
	}
	resp, err := utils.GetUserSvrClient().GetUserInfo(ctx, &pb.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		log.Errorf("get user info error: %v", err)
		response.Fail(ctx, err.Error(), nil)
		return
	}
	log.Infof("get user info: %+v", resp)
	response.Success(ctx, "success", resp.UserInfo)
}

func UserRegister(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	if len(username) > 32 || len(password) > 32 {
		response.Fail(ctx, "username or password invalid", nil)
		return
	}
	log.Info(username, password)
	resp, err := utils.GetUserSvrClient().Register(ctx, &pb.RegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		zap.L().Error("register error", zap.Error(err))
		response.Fail(ctx, err.Error(), nil)
		return
	}
	response.Success(ctx, "success", resp)
}
