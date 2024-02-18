package routes

import (
	"gatewaysvr/controller"
	"gatewaysvr/utils/middleWare"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup) {
	user := r.Group("user")
	{
		user.GET("/", middleWare.AuthMiddleWare(), controller.GetUserInfo)
		user.POST("/login/", controller.UserLogin)
		user.POST("/register/", controller.UserRegister)
	}
}
