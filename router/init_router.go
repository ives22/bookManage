package router

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由并返回一个 *gin.Engine 实例
func InitRouter() *gin.Engine {
	r := gin.Default() // 创建默认的 Gin 引擎实例
	TestRouters(r)     // 注册测试路由
	SetupApiRouters(r) // 注册 API 路由
	return r           // 返回初始化后的 Gin 引擎实例
}