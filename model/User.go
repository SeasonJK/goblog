package model

import (
	"goblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validator:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validator:"required,min=6,max=120" label:"密码"`
	Role int `gorm:"type:int;DEFAULT:2" json:"role" validator:"required,get=2" label:"角色码"`
}

// 查询用户是否存在
func CheckUser(name string)(code int){
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	// 查询用户id>0，则说明用户已存在，返回code码
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 更新查询
func UpdateUser(id int, name string)(code int){
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id){
		return errmsg.SUCCESS
	}
	if user.ID > 0{
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	err = db.Create(data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 查询用户
func GetUser(id int)(User, int){
	var user User
	err = db.Where("ID = ?", id).First(&user).Error
	if err != nil{
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	return user, errmsg.SUCCESS
}
// 查询用户列表(分页显示)
func GetUsers(username string, pageSize, pageNum int)([]User, int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id, username, role").Where(
			"username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
		db.Model(&users).Where("username LIKE ?", username+"%").Count(&total)
		return users, total
	}
	db.Select("id, username, role").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	db.Model(&users).Count(&total)
	if err != nil{
		return users, 0
	}
	return users, total
}
// 编辑用户信息
func EditUser(id int, data *User)int{
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 修改密码
func ChangePassword(id int, data *User)int{
	//var user User
	//var maps = make(map[string]interface{})
	//maps["password"] = data.Password
	err = db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 删除用户
func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 生成密码
func ScryptPassword(password string)string{
	const cost = 10

	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil{
		log.Fatal(err)
	}
	return string(HashPw)
}

// 密码加密
func (u *User)BeforeSave(_ *gorm.DB)(err error){
	u.Password = ScryptPassword(u.Password)
	return nil
}
func (u *User)BeforeUpdate(_ *gorm.DB)(err error){
	u.Password = ScryptPassword(u.Password)
	return nil
}

//todo 后台登录验证

//todo 前台登录