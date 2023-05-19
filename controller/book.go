package controller

import (
	"github.com/gin-gonic/gin"
	"myBookManage/dao/mysql"
	"myBookManage/model"
	"net/http"
	"strconv"
)

// CreateBookHandler 处理创建书籍的请求
func CreateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	// 判断是否有传入用户
	if len(p.Users) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  "missing parameter for users",
		})
		return
	}

	// 如果有用户，需判断用户是否已存在（这里不做复杂的用户校验，就是只是判断传入的用户是否存在，如果不存在，则不允许创建书籍）
	var users []model.User
	existingUsers := make(map[int64]struct{}) // 用来校验是否添加到users中
	for _, v := range p.Users {
		var user model.User
		if rows := mysql.DB.Debug().Where("username = ?", v.Username).Or("uid = ?", v.UID).First(&user).Row(); rows != nil {
			// 如果传入的用户存在，则将这个用户记录添加到 users 切片中
			// 首先判断 users 切片中是否已经存在该用户了，如果存在则不添加
			if _, ok := existingUsers[user.UID]; !ok {
				users = append(users, user)
				existingUsers[user.UID] = struct{}{}
			}
		}
	}

	// 添加书籍
	book := model.Book{
		Name: p.Name,
		Desc: p.Desc,
	}
	mysql.DB.Create(&book)

	// 为用户和书籍添加关联关系
	mysql.DB.Model(&book).Association("Users").Append(users)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "Add book success"})
}

// GetBookListHandler 处理获取所有书籍及其关联用户的请求
func GetBookListHandler(c *gin.Context) {
	books := []model.Book{}
	// 查看所有书籍，及其相关的用户
	mysql.DB.Preload("Users").Find(&books)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"books": books,
		},
	})
}

// GetBookDetailHandler 查看指定的书籍
func GetBookDetailHandler(c *gin.Context) {
	pipelineIdStr := c.Param("id") // 获取URL参数
	bookID, _ := strconv.ParseInt(pipelineIdStr, 10, 64)
	book := model.Book{ID: bookID}
	// 查询书籍，及其关联的用户
	mysql.DB.Preload("Users").Find(&book)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"book": book,
		},
	})
}

// UpdateBookHandler 修改书籍信息
func UpdateBookHandler(c *gin.Context) {
	p := new(model.Book)
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	//	判断传入的书籍是否存在，如果存在则进行修改相关的操作
	oldBook := model.Book{ID: p.ID}
	if rows := mysql.DB.Debug().Where(&oldBook).First(&oldBook).Row(); rows != nil {
		var newBook model.Book
		if p.Name != "" {
			newBook.Name = p.Name
		}
		if p.Desc != "" {
			newBook.Desc = p.Desc
		}

		// 判断是否有传入用户信息，如果传入了用户，则还要修改书籍表和用户的关系
		if len(p.Users) != 0 {
			// 首先获取用户传入的用户信息是否存在。
			var users []model.User
			existingUsers := make(map[int64]struct{}) // 用来校验是否添加到users中
			for _, v := range p.Users {
				var user model.User
				if rows := mysql.DB.Debug().Where("username = ?", v.Username).Or("uid = ?", v.UID).First(&user).Row(); rows != nil {
					// 如果传入的用户存在，则将这个用户记录添加到 users 切片中
					// 首先判断 users 切片中是否已经存在该用户了，如果存在则不添加
					if _, ok := existingUsers[user.UID]; !ok {
						users = append(users, user)
						existingUsers[user.UID] = struct{}{}
					}
				} else {
					//	如果传入的用户再数据库中没有找到，那么不允许更新。因为参数错误
					c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "user does not exist"})
					return
				}
			}
			// 修改书籍信息
			mysql.DB.Model(&oldBook).Updates(newBook)
			// 修改书籍和用户之间的绑定关系
			// 这里用 Replace 替换测试过，没有效果，所以使用下面的删除关联，然后再添加关联。
			//mysql.DB.Debug().Model(&oldBook).Association("Users").Replace(users)
			// 先删除之前的关联，然后重新关联
			mysql.DB.Debug().Model(&oldBook).Association("Users").Clear()
			mysql.DB.Model(&oldBook).Association("Users").Append(users)
		} else {
			// 如果没有传入用户细腻些，只是修改修改书籍信息
			mysql.DB.Model(&oldBook).Updates(newBook)
		}
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "update book success"})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"code": 404, "msg": "book does not exist"})
	}
}

// DeleteBookHandler 删除书籍信息
func DeleteBookHandler(c *gin.Context) {
	pipelineIdStr := c.Param("id") // 获取URL参数
	bookID, _ := strconv.ParseInt(pipelineIdStr, 10, 64)
	// 删除book时，同时删除第三张表中的关联信息，也就是与用户对应关系记录
	mysql.DB.Debug().Select("Users").Delete(&model.Book{ID: bookID})
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}