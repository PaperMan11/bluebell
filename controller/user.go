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
		"msg": "注册成功",
	})
}

// 用户登录
func LoginHandler(ctx *gin.Context) {
	// 获取参数校验
	// 获取参数和参数校验
	p := new(models.ParamLogin)
	if err := ctx.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	// 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "用户或密码错误",
		})
		return
	}
	// 返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}
