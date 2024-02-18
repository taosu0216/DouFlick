package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"
	"github.com/gin-gonic/gin"
)

func VideoRouter(r *gin.RouterGroup) {
	video := r.Group("publish")
	{
		video.POST("/action/", middleWare.AuthMiddleWare(), controller.Publish)
		video.GET("/list/", middleWare.AuthMiddleWare(), controller.GetVideoList)
	}
}
