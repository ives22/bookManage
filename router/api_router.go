package router

import (
	"github.com/gin-gonic/gin"
	"myBookManage/controller"
	"myBookManage/middleware"
)

// SetupApiRouters 设置 API 路由
func SetupApiRouters(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler) // 注册用户的路由
	r.POST("/login", controller.LoginHandler)       // 用户登录的路由

	v1 := r.Group("/api/v1")
	v1.Use(middleware.AuthMiddleware()) // 给v1这个路由组注册auth中间件

	v1.POST("book", controller.CreateBookHandler)       // 创建书籍的路由
	v1.GET("book", controller.GetBookListHandler)       // 获取书籍列表的路由
	v1.GET("book/:id", controller.GetBookDetailHandler) // 获取单个书籍详情的路由
	v1.PUT("book", controller.UpdateBookHandler)        // 更新书籍信息的路由
	v1.DELETE("book/:id", controller.DeleteBookHandler) // 删除书籍的路由
}