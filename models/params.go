package models

// 请求参数
// binding:"required" 请求参数不能为空
// binding:"required,eqfield=PassWord" 请求参数不能为空且字段必须等于 PassWord
type ParamSignUp struct {
	UserName   string `json:"username" binding:"required"`
	PassWord   string `json:"password" binding:"required"`
	RePassWord string `json:"re_password" binding:"required,eqfield=PassWord"`
}
