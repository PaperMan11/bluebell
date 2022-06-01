package mysql

import (
	"bluebell/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

// CreatePost 创建帖子
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostByID 根据id查询单个帖子的数据
func GetPostByID(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(data, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(offset int64, limit int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2)
	db.Select(&posts, sqlStr, offset-1, limit)
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据（结果顺寻为给定的顺序）
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
