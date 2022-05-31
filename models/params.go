package models

// 请求参数
// binding:"required" 请求参数不能为空
// binding:"required,eqfield=PassWord" 请求参数不能为空且字段必须等于 PassWord
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"re_password" binding:"required,eqfield=PassWord"`
}

// 登录请求参数
type ParamLogin struct {
	UserName string `json:"username" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}

// 投票数据
type ParamVoteData struct {
	// UserID // 从请求中获取当前用户
	// 帖子id
	PostID string `json:"post_id" binding:"required"`
	// 赞成票（1）反对票（-1）取消投票（0）
	// binding:"required,oneof=1 -1 0" 参数不能为空且值为其中一个
	Direction int8 `json:"direction,string" binding:"oneof=1 0 -1"`
}
