package goblog

import (
	"goblog/model"
	"goblog/routes"
)

func main(){
	// 引用数据库
	model.InitDatabase()
	// 引用路由组件
	routes.InitRouter()
}