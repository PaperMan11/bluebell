package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	// 生成 post.id
	p.ID = snowflake.GetID()
	// 保存到数据库
	if err = mysql.CreatePost(p); err != nil {
		return err
	}
	redis.CreatePost(p.ID)
	return
}

func GetPostByID(pid int64) (data *models.ApiPostDetail, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Error(err))
		return
	}
	// 根据作者 id 查询作者信息
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(pid) failed", zap.Error(err))
		return
	}
	// 根据社区 id 查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
		return
	}

	// 数据拼接
	data = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(offset int64, limit int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(offset, limit)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者 id 查询作者信息
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(pid) failed", zap.Error(err))
			continue
		}
		// 根据社区 id 查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID() failed", zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
