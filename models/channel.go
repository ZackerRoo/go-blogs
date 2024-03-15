package models

// 数据库，用户数据映射模型
type Channel struct {
	Id      uint64 `form:"id" gorm:"primaryKey"`
	Title   string `form:"title" gorm:"title"`
	Slug    string `form:"slug" gorm:"slug"`
	Content string `form:"content" gorm:"content"`
	Status  int    `form:"status" gorm:"status"`
	Weight  int    `form:"weight" gorm:"weight"`
}
