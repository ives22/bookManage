package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestRouters(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.GET("test", TestHandler)

}

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}