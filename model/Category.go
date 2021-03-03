package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

//todo 查询分类是否存在

//todo	新增分类

//todo	查询单个分类信息

//todo	查询分类列表

//todo	编辑分类信息

//todo	删除分类