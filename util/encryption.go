package util

import (
	"crypto/md5"
	"fmt"
)

// GetMd5 获取字符串md5
func GetMd5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
