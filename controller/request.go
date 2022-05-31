package controller

import (
	"errors"
	"strconv"

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

// getPageInfo 获取分页参数
func getPageInfo(ctx *gin.Context) (int64, int64) {
	offsetStr := ctx.Query("offset")
	limitStr := ctx.Query("limit")
	var (
		limit  int64
		offset int64
		err    error
	)
	offset, err = strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		offset = 1
	}
	limit, err = strconv.ParseInt(limitStr, 10, 64)
	if err != nil {
		limit = 10
	}
	return offset, limit
}
