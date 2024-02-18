package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"
	"github.com/gin-gonic/gin"
)

func FavoriteRouter(r *gin.RouterGroup) {
	favorite := r.Group("favorite")
	{
		favorite.POST("/action/", middleWare.AuthMiddleWare(), controller.FavoriteAction)
		favorite.GET("/list/", middleWare.AuthWithOutMiddleware(), controller.GetFavoriteList)
	}
}
