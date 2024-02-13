package routes

import (
	"gatewaysvr/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.GET("/", controller.GetUserInfo)
		user.POST("/login/", controller.UserLogin)
	}
}
