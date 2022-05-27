package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

// 业务逻辑处理

func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.UserName); err != nil {
		// 数据库查询出错
		return err
	}

	// 生成uid
	userID := snowflake.GetID()
	// 构造一个 Uesr 实例
	user := &models.User{
		UserID:   userID,
		Username: p.UserName,
		Password: p.PassWord,
	}
	// 保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (err error) {
	user := &models.User{
		Username: p.UserName,
		Password: p.PassWord,
	}
	return mysql.Login(user)
}
