package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	// 查询数据库，查找所有的 community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	// 查询数据库，查找指定 id 的 community 并返回
	return mysql.GetCommunityDetailByID(id)
}
