package pkg

import "golang.org/x/crypto/bcrypt"

// HashPassword 生成密码的哈希值
// 使用 bcrypt 算法将原始密码进行哈希，并返回哈希值的字符串表示和可能的错误
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPasswordHash 验证密码的哈希值
// 使用 bcrypt 算法将原始密码与哈希值进行比较，判断密码是否匹配
// 如果密码匹配，则返回 true，否则返回 false
func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}