package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"

	"github.com/gin-gonic/gin"
)

func CommentRouter(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.GET("/list", middleWare.AuthWithOutMiddleware(), controller.GetCommentList)
		comment.POST("/action", middleWare.AuthMiddleWare(), controller.AddComment)
	}
}
