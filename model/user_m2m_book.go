package model

type BookUser struct {
	ID      uint64 `gorm:"primaryKey"`         // 连接表的唯一标识符，作为主键
	UserUID uint64 `gorm:"primaryKey"`         // 用户的唯一标识符
	BookID  uint64 `gorm:"primaryKey"`         // 书籍的唯一标识符
	User    User   `gorm:"foreignKey:UserUID"` // 外键关联到 User 模型的 UserID 字段
	Book    Book   `gorm:"foreignKey:BookID"`  // 外键关联到 Book 模型的 BookID 字段
}

func (BookUser) TableName() string {
	return "book_users" // 指定表名为 "book_users"
}