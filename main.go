package main

import (
	"myBookManage/dao/mysql"
	"myBookManage/router"
)

func main() {
	// 初始化mysql连接
	mysql.InitMysql()

	//	初始化路由
	r := router.InitRouter()

	// 运行服务器，监听端口 ":8888"
	r.Run(":8888")
}