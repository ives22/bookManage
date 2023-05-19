package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"myBookManage/dao/mysql"
	"myBookManage/model"
	"net/http"
)

func AuthMiddleware() func(c *gin.Context) {
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// token验证成功，返回c.Next继续，否则返回c.Abort()直接返回
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		var u model.User
		//	如果没有当前用户
		fmt.Println("token")
		// created_at != update_at 查询语句加上这个，是为了避免注册了的用户从未登录过token也为空的情况。 因为一旦登录过，那么就会生成token，那么update_at就会发生变换
		if rows := mysql.DB.Debug().Where("token = ? AND created_at != update_at", token).First(&u).RowsAffected; rows != 1 {
			fmt.Println("来了")
			c.JSON(http.StatusForbidden, gin.H{"code": 403, "msg": "Authorization failed"})
			c.Abort()
			return
		}
		fmt.Println("没来")
		//	将当前请求的userID信息保存到请求的上下午c上
		c.Set("UserId", u.UID)
		c.Next()
	}
}