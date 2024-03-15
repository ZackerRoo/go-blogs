package service

import (
	"go-blogs/common/global"
	"go-blogs/models"
)

type PostService struct {
}

// 添加
func (p *PostService) AddPost(post models.Post) int64 {
	return global.Db.Table("post").Create(&post).RowsAffected
}

// 删除
func (p *PostService) DelPost(id int) int64 {
	return global.Db.Delete(&models.Post{}, id).RowsAffected
}

// 修改
func (p *PostService) UpdatePost(post models.Post) int64 {
	return global.Db.Updates(&post).RowsAffected
}

// get
func (p *PostService) GetPost(id int) models.Post {
	var post models.Post
	global.Db.First(&post, id)
	return post
}

// get post list
func (p *PostService) GetPostList() []models.Post {
	postList := make([]models.Post, 0)
	global.Db.Find(&postList)
	return postList
}
