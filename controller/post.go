package controller

import (
	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreatePostHandler(ctx *gin.Context) {
	// 获取参数及参数校验
	p := new(models.Post)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param")
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	// 从 ctx 中获取到当前请求用户的ID
	userID, err := getCurrentUserID(ctx)
	if err != nil { // 用户需要登录
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	p.AuthorID = userID
	// 创建帖子（存入数据库）
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, nil)
}
