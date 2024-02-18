package routes

import (
	"gatewaysvr/config"
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"
	"github.com/gin-gonic/gin"
)

func RouteInit() *gin.Engine {
	if config.GetGlobalConfig().SvrConfig.Mode == gin.ReleaseMode {
		// gin设置成发布模式：gin不在终端输出日志
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	douyin := r.Group("/douyin")
	{
		UserRouter(douyin)
		CommentRouter(douyin)
		//TODO:以下内容
		VideoRouter(douyin)
		FavoriteRouter(douyin)
		RelationRouter(douyin)
		douyin.GET("/feed/", middleWare.AuthWithOutMiddleware(), controller.Feed)
	}
	return r
}
