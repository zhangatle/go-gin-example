package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangatle/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "hello go",
		})
	})
	return r
}
