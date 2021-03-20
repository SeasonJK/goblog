package model

import "goblog/utils/errmsg"

type Profile struct {
	ID 		int `gorm:"primaryKey" json:"id"`
	Name 	string `gorm:"type:varchar(20)" json:"name"`
	Desc 	string `gorm:"tyep:varchar(200)" json:"desc"`
	QQChat 	string `gorm:"tyep:varchar(200)" json:"qq_chat"`
	WeChat 	string `gorm:"tyep:varchar(100)" json:"we_chat"`
	WeiBo	string `gorm:"tyep:varchar(200)" json:"wei_bo"`
	Bili	string `gorm:"tyep:varchar(200)" json:"bili"`
	Email 	string `gorm:"tyep:varchar(200)" json:"email"`
	Img 	string `gorm:"tyep:varchar(200)" json:"img"`
	Avatar	string `gorm:"tyep:varchar(200)" json:"avatar"`
	IcpRecord	string `gorm:"tyep:varchar(200)" json:"icp_record"`
}

// 获取个人信息配置
func GetProfile(id int)(Profile, int){
	var profile Profile
	err = db.Where("ID = ?", id).First(&profile).Error
	if err != nil{
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

// 更新个人信息设置
func UpdateProfile(id int, data *Profile)int{
	var profile Profile
	err = db.Model(&profile).Where("ID = ?", id).Updates(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
