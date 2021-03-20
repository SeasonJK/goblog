package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 获取个人信息配置
func GetProfile(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}

// 更新个人设置
func UpdateProfile(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Profile
	_ = c.ShouldBindJSON(&data)
	code := model.UpdateProfile(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message": 	errmsg.GetErrMsg(code),
	})
}