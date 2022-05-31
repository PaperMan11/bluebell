package mysql

import "bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
		post_id, title, content, author_id, community_id)
		values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(pid int64) (data *models.Post, err error) {
	data = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(data, sqlStr, pid)
	return
}
