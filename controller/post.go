package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreatePostHandler 创建帖子
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

// GetPostDetailHandler 获取帖子详情的处理函数
func GetPostDetailHandler(ctx *gin.Context) {
	// 获取参数（从URL中获取帖子的id）
	pidStr := ctx.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	// 根据 id 取出帖子数据（查数据库）
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostByID(pid) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}

// GetPostListHandler 获取帖子列表的处理函数
func GetPostListHandler(ctx *gin.Context) {
	// 获取分页参数
	offset, limit := getPageInfo(ctx)
	// 获取数据
	data, err := logic.GetPostList(offset, limit)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}

// 升级版帖子列表接口
// 根据前端传过来的参数（按分数、按创建时间排序）动态的去获取帖子的列表
// 1、获取参数
// 2、去 redis 查询 id 列表
// 3、根据 id 去数据库查询帖子详细信息
func GetPostListHandler2(ctx *gin.Context) {
	// 获取分页参数（GET请求参数 /api/v1/postlsit2?offset=1&limit=2）
	p := &models.ParamPostList{
		Offset: 1,
		Limit:  10,
		Order:  models.OrderTime,
	}
	if err := ctx.ShouldBindQuery(p); err != nil {
		zap.L().Error("GetPostListHandler2 with invalid param", zap.Error(err))
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	// 获取数据
	data, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(ctx, data)
}
