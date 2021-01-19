package mysql

import "bluebell/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post
	(post_id,title,content,author_id, community_id)
	values(?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return err
}
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title , content, author_id , create_time , community_id
               from post where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	return post, err
}
func GetPostList(page int64, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title , content, author_id , create_time , community_id
               from post  limit ?, ?`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
