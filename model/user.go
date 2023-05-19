package model

import (
	"gorm.io/gorm"
	"myBookManage/pkg"
	"time"
)

type User struct {
	UID       int64     `gorm:"primary_key" json:"uid"`                                                                 // 用户ID，主键
	Username  string    `gorm:"not null; type:varchar(128); unique" json:"username" binding:"required"`                 // 用户名，非空、唯一
	Password  string    `gorm:"not null; type:varchar(128)" json:"password" binding:"required"`                         // 密码，非空
	Mobile    string    `gorm:"unique; type:varchar(32)" json:"mobile"`                                                 // 手机号，唯一
	Email     string    `gorm:"unique; type:varchar(32)" json:"email"`                                                  // 邮箱，唯一
	Gender    int64     `json:"gender"`                                                                                 // 性别，0表示男 1表示女
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at"`                   // 创建时间，使用当前时间作为默认值，仅在创建时设置
	UpdateAt  time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"update_at"` // 更新时间，使用当前时间作为默认值，并在更新时自动更新
	Token     string    `gorm:"type:varchar(128)" json:"token"`                                                         // token
	Books     []Book    `gorm:"many2many:book_users"`                                                                   // 与书籍之间的多对多关系
}

func (User) TableName() string {
	return "user" // 指定表名为 "user"
}

// BeforeCreate 利用GORM的钩子函数，往数据库中存入数据的时候，进行加密处理
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bcryptPW, err := pkg.HashPassword(u.Password) // 使用密码哈希函数对用户密码进行哈希处理
	if err != nil {
		return err // 如果密码哈希处理出现错误，则直接返回错误
	}
	u.Password = bcryptPW // 将哈希后的密码赋值给用户密码字段
	return nil            // 返回 nil 表示没有错误发生
}