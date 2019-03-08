package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// md5 加密
func EncryptStr(str string) string {
	salt := "lincktek"
	pattern := []byte(str + salt)
	has := md5.New()
	code := has.Sum(pattern)
	return hex.EncodeToString(code)
}

