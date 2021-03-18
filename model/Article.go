package model

import (
	"goblog/utils/errmsg"
	"gorm.io/gorm"
)

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

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
// 查询分类下的所有文章
func GetCateArt(id int, pageSize, pageNum int) ([]Article, int64, int) {
	var CateArtList []Article
	var total int64

	err := db.Select("article.id,title, `cid`, img, created_at, `desc`, comment_count, read_count").Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&CateArtList).Error
	db.Model(&CateArtList).Where("cid = ?", id).Count(&total)

	if err != nil{
		return CateArtList, 0, errmsg.ERROR_CATE_NOT_EXIST
	}
	return CateArtList, total, errmsg.SUCCESS
}
// 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var article Article
	err := db.Where("id = ?", id).Preload("Category").First(&article).Error
	db.Model(&article).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))

	if err != nil{
		return article, errmsg.ERROR_ART_NOT_EXIST
	}
	return article, errmsg.SUCCESS
}

// 查询文章列表
func GetArtList(title string, pageSize,pageNum int)([]Article, int, int64){
	var articleList []Article
	var err error
	var total int64

	err = db.Select("article.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}

// 查询文章标题
func SearchArtTitle(title string, pageSize,pageNum int)([]Article, int, int64){
	var articleList []Article
	var err error
	var total int64
	err = db.Select("article.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, category.name").Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
		title+"%").Find(&articleList).Count(&total).Error
	if err != nil{
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, total
}
// 编辑文章
func EditArticle(id int, data *Article) int {
	var article Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["description"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Model(&article).Where("id = ?", id).Updates(&maps).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err = db.Where("id = ?", id).Delete(&article).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}






















