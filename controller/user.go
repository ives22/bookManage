package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"myBookManage/dao/mysql"
	"myBookManage/model"
	"myBookManage/pkg"
	"net/http"
)

// RegisterHandler 处理用户注册请求
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	mysql.DB.Create(&p) // 创建用户记录

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Register user success",
	})
}

// LoginHandler 处理用户登录请求
func LoginHandler(c *gin.Context) {
	p := new(model.User)
	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// 去数据库中查找该用户，如果存在，则继续进行下面的密码验证
	u := model.User{Username: p.Username}
	if rows := mysql.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "user does not exist",
		})
		return
	}

	// 对密码进行校验
	if !pkg.CheckPasswordHash(u.Password, p.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 401,
			"msg":  "Password error",
		})
		return
	}

	token := uuid.New().String()              // 生成新的token
	mysql.DB.Model(&u).Update("token", token) // 更新用户的令牌字段
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "Login success",
		"data": gin.H{"token": token},
	})
}