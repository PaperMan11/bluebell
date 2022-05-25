package routes

import (
	"bluebell/logger"
	"bluebell/settings"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, settings.Conf.Version)
	})
	return r
}
