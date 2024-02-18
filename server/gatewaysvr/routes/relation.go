package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"
	"github.com/gin-gonic/gin"
)

func RelationRouter(r *gin.RouterGroup) {
	relation := r.Group("relation")
	{
		relation.POST("/action/", middleWare.AuthMiddleWare(), controller.RelationAction)
		relation.GET("/follow/list/", middleWare.AuthWithOutMiddleware(), controller.GetFollowList)
		relation.GET("/follower/list/", middleWare.AuthWithOutMiddleware(), controller.GetFollowerList)
	}
}
