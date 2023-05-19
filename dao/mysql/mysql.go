package mysql

import (
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"myBookManage/model"
)

var DB *gorm.DB

// InitMysql 初始化 MySQL 数据库连接
func InitMysql() {
	// 定义数据库连接字符串
	dsn := "root:admin123@tcp(124.71.33.240:3306)/book_manage?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开数据库连接
	db, err := gorm.Open(gmysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 将数据库连接赋值给全局变量
	DB = db

	//	自动创建表结构
	if err := DB.AutoMigrate(&model.User{}, &model.Book{}, &model.BookUser{}); err != nil {
		panic(err)
	}
}