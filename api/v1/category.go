package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

//	添加分类
func AddCategory(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS{
		model.CreateCategory(&data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		code = errmsg.ERROR_CATENAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})
}
//	查询分类信息
func GetCateInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetCateInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"message":	errmsg.GetErrMsg(code),
	})

}
//	查询分类列表
func GetCateList(c *gin.Context){
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0{
		pageNum = 1
	}

	data, total := model.GetCateList(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"total":	total,
		"message":	errmsg.GetErrMsg(code),
	})
}
//	编辑分类名
func EditCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS{
		model.EditCate(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED{
		c.Abort()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data": 	data,
		"message":	errmsg.GetErrMsg(code),
	})
}
//	删除分类
func DeleteCategory(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),
	})
}




















