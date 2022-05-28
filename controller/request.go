package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// ctx: key
const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUserID 获取当前登录的 userID
func getCurrentUserID(ctx *gin.Context) (userID int64, err error) {
	uid, ok := ctx.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
