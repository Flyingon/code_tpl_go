package util

import (
	"reflect"
	"unsafe"
)

// BytesToStringFast convert bytes to string quickly by avoiding underlying (indirect part) memory allocation
func BytesToStringFast(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringToBytesFast convert string to []byte quickly by avoiding underlying (indirect part) memory allocation
func StringToBytesFast(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

// Find return the index of string `x` in []string `a`
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return -1
}

func Substr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

// IsStrInList 判断字符串e是否在字符串数组l中
func IsStrInList(e string, l []string) bool {
	for _, v := range l {
		if e == v {
			return true
		}
	}
	return false
}
