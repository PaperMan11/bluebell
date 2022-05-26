package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, settings.Conf.Version)
	})
	return r
}
