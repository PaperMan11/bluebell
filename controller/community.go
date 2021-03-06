package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/* 跟社区相关 */
func CommunityHandler(ctx *gin.Context) {
	// 查询到所有的社区（community_id, community_name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(ctx *gin.Context) {
	// 从 url上 获取社区 id
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(ctx, CodeInvalidParam)
		return
	}

	// 查询到所有的社区（community_id, community_name）以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, data)
}
