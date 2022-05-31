package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostVoteHandler(ctx *gin.Context) {
	// 参数校验
	p := new(models.ParamVoteData)
	if err := ctx.ShouldBindJSON(p); err != nil {
		fmt.Println(err)
		ResponseError(ctx, CodeInvalidParam)
		return
	}
	userID, err := getCurrentUserID(ctx)
	if err != nil {
		ResponseError(ctx, CodeNeedLogin)
		return
	}
	// 具体投票的业务逻辑
	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed", zap.Error(err))
		ResponseError(ctx, CodeServerBusy)
		return
	}
	ResponseSuccess(ctx, nil)
}
