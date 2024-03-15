package models

import "gorm.io/gorm"

// 数据库，博客数据映射模型
type Post struct {
	gorm.Model
	Title     string `gorm:"slug"`
	Thumbnail string `gorm:"thumbnail"`
	Summary   string `gorm:"summary"`
	Content   string `gorm:"content"`
	AuthorId  int    `gorm:"author_id"`
	ChannelId int    `gorm:"channel_id"`
	Comments  int    `gorm:"comments"`
	Favors    int    `gorm:"favors"`
	Featured  int    `gorm:"featured"`
	Status    int    `gorm:"status"`
	Tags      string `gorm:"tags"`
	Views     int    `gorm:"views"`
	Weight    int    `gorm:"weight"`
	Url       string `gorm:"url"`
}
