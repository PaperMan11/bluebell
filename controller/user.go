package controller

import (
	"fmt"
	"net/http"

	"bluebell/logic"
	"bluebell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 处理注册请求
func SignUpHandler(ctx *gin.Context) {
	// 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(*p)

	// 业务处理
	if err := logic.SignUp(p); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}

	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
