package model

import (
	"fmt"
	"goblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB
var err error

func InitDatabase(){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.Dbname)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm 日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 单数表名，
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
	}

	// 自动迁移
	_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDb, _ := db.DB()
	// 设置连接池最大的连接数量
	sqlDb.SetMaxIdleConns(10)

	// 设置数据库的最大连接数量
	sqlDb.SetMaxOpenConns(100)

	// 设置连接的最大可复用时间
	sqlDb.SetConnMaxLifetime(10 * time.Second)


}
