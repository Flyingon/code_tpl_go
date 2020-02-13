package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

// GetMd5V1 获取字符串md5
func GetMd5V1(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetMd5V2 获取字符串md5
func GetMd5V2(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

// GetMd5V3 获取字符串md5
func GetMd5V3(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
