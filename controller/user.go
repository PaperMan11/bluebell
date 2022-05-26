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
	var p models.ParamSignUp
	if err := ctx.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	// 手动对请求参数进行详细业务要求校验
	if len(p.UserName) == 0 || len(p.PassWord) == 0 || len(p.RePassWord) == 0 || p.PassWord != p.RePassWord {
		zap.L().Error("SignUp with invalid param")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "请求参数有误",
		})
		return
	}
	fmt.Println(p)
	// 业务处理
	logic.SignUp()
	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
