package model

import "time"

type Book struct {
	ID        int64     `gorm:"primary_key" json:"id"`                                                                  // 书籍ID，主键
	Name      string    `gorm:"not null; type:varchar(128); unique" json:"name" binding:"required"`                     // 书籍名称，非空、唯一
	Desc      string    `gorm:"type:varchar(256)" json:"desc"`                                                          // 书籍描述
	CreatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP;<-:create" json:"created_at"`                   // 创建时间，使用当前时间作为默认值，仅在创建时设置
	UpdatedAt time.Time `gorm:"type:TIMESTAMP;default:CURRENT_TIMESTAMP  on update current_timestamp" json:"update_at"` // 更新时间，使用当前时间作为默认值，并在更新时自动更新
	Users     []User    `gorm:"many2many:book_users"`                                                                   // 与用户之间的多对多关系
}

func (Book) TableName() string {
	return "book" // 指定表名为 "book"
}