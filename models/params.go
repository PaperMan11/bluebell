package models

// 请求参数
type ParamSignUp struct {
	UserName   string `json:"username"`
	PassWord   string `json:"password"`
	RePassWord string `json:"re_password"`
}
