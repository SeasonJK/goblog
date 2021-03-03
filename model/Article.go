package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title 			string	`gorm:"type:varchar(100);not null" json:"title"`
	Cid 			int	`gorm:"type:int;not null" json:"cid"`
	Desc 			string	`gorm:"type:varchar(200)" json:"desc"`
	Content 		string	`gorm:"type:longtext" json:"content"`
	Img 			string	`gorm:"type:varchar(100)" json:"img"`
	CommentCount	int	`gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount		int	`gorm:"type:int;not null;default:0" json:"read_count"`
}

//todo 新增文章

//todo 查询分类下的所有文章

//todo 查询单个文章

//todo 查询文章列表

//todo 编辑文章

//todo 删除文章
