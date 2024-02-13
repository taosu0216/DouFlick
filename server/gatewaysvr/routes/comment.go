package routes

import (
	"gatewaysvr/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoute(r *gin.RouterGroup) {
	comment := r.Group("comment")
	{
		comment.GET("/list", controller.GetCommentList)
		comment.POST("/action", controller.AddComment)
	}
}
