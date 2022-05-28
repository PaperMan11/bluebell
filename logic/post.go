package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	// 生成 post.id
	p.ID = snowflake.GetID()
	// 保存到数据库
	return mysql.CreatePost(p)
}
