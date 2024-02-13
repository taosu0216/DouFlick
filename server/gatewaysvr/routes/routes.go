package routes

import "github.com/gin-gonic/gin"

func RouteInit() *gin.Engine {
	r := gin.New()
	douyin := r.Group("/douyin")
	{
		UserRoute(douyin)
	}
	return r
}
