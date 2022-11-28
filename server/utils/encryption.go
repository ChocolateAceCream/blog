package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD-5 encoding
func MD5Encryption(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	r := hex.EncodeToString(s.Sum(nil))
	return strings.ToUpper(r)
}
