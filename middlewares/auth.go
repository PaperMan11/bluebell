package middlewares

import (
	"bluebell/controller"
	"bluebell/pkg/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于 JWT 的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 客户端携带 Token 有三种方式 1、放在请求头中 2、放在请求体 3、放在URI
		// 这里假设 Token 放在 Header 的 Authorization 中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		// Postman: Bearer Token
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			controller.ResponseError(ctx, controller.CodeNeedLogin)
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(ctx, controller.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// parts[1] 是获取到 tokenString
		// 解析 token -> *jwt.MyClaims
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(ctx, controller.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// 将当前请求的 username 信息保存到请求的上下文 ctx 上
		ctx.Set(controller.CtxUserIDKey, mc.UserID)
		// 后续处理请求的函数中，可以通过 ctx.Get(CtxUserIDKey) 来获取当前请求的用户信息
		ctx.Next()
	}
}
