package logic

import (
	"bluebell/dao/mysql"
	"bluebell/pkg/snowflake"
)

// 业务逻辑处理

func SignUp() {
	// 判断用户是否存在
	mysql.QueryUserByName()
	// 生成uid
	snowflake.GetID()
	// 保存进数据库
	mysql.InsertUser()
}
