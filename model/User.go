package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,get=2" label:"角色码"`
}

//todo 查询用户是否存在


//todo 新增用户

//todo 查询用户

//todo 查询用户列表

//todo 编辑用户信息

//todo 修改密码

//todo 删除用户

//todo 密码加密

//todo 生成密码

//todo 后台登录验证

//todo 前台登录