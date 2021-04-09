package v1

import (
	"github.com/gin-gonic/gin"
	"goblog/model"
	"goblog/utils/errmsg"
	"net/http"
	"strconv"
)

// 新增文章
func AddArticle(c *gin.Context){
	var data model.Article
	_ = c.ShouldBindJSON(data)

	code := model.CreateArticle(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}
// 查询分类下所有文章
func GetCateArt(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	//分页器
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code :=model.GetCateArt(id, pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}
// 查询单个文章信息
func GetArtInfo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	data, code := model.GetArtInfo(id)

	c.JSON(http.StatusOK, gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArtList(c *gin.Context){
	title := c.Query("title")
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	//分页器
	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	if pageNum == 0 {
		pageNum = 1
	}

	if len(title) == 0{
		data, code, total := model.GetArtList(title, pageSize, pageNum)
		c.JSON(http.StatusOK, gin.H{
			"status":		code,
			"data":			data,
			"total":		total,
			"message":		errmsg.GetErrMsg(code),
		})
		return
	}

	data, code, total := model.SearchArtTitle(title, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"data":		data,
		"total":	total,
		"message":	errmsg.GetErrMsg(code),
	})
}
// 编辑文章
func EditArticle(c *gin.Context){
	var data model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.EditArticle(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message": 	errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":	code,
		"message":	errmsg.GetErrMsg(code),
	})
}
























