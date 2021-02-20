package util

import (
	"reflect"
	"strings"
	"unicode/utf8"
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

// Setw 用c补充s长度到l, 类似C++: std::setw(16) << std::setfill('0')
func Setw(l uint, s, c string) (r string) {
	if len(c) > 1 {
		return s
	}
	var num uint
	if uint(len(s)) < l {
		num = l - uint(len(s))
	}
	for i := uint(0); i < num; i++ {
		r += c
	}
	r += s
	return r
}

// SubStrDecodeRuneInString 基于UTF8编码截取字段
func SubStrDecodeRuneInString(s string, length int) string {
	var size, n int
	for i := 0; i < length && n < len(s); i++ {
		_, size = utf8.DecodeRuneInString(s[n:])
		n += size
	}

	return s[:n]
}

// keyAddElem 添加元素
// 输入: curVal: "a,b,c" addVal: "d,e" splitKey: ","
// 返回: a,b,c -> a,b,c,d,e
func keyAddElems(curVal string, addVal string, splitKey string) string {
	newVal := curVal
	if curVal != "" && addVal != "" {
		oldValList := strings.Split(curVal, splitKey)
		addValList := strings.Split(addVal, splitKey)
		oldValMap := make(map[string]bool)
		for _, oldKey := range oldValList {
			oldValMap[oldKey] = true
		}
		for _, add := range addValList {
			if !oldValMap[add] {
				oldValList = append(oldValList, add)
			}
		}
		newVal = strings.Join(oldValList, splitKey)
	} else if curVal == "" {
		newVal = addVal
	}
	return newVal
}

// keyDelElems 添加元素
// 输入: curVal: "a,b,c,d,e" delVal: "a,e" splitKey: ","
// 返回: b,c,d
func keyDelElems(curVal string, delVal string, splitKey string) string {
	newVal := curVal
	if curVal != "" && delVal != "" {
		oldValList := strings.Split(curVal, splitKey)
		delValList := strings.Split(delVal, splitKey)
		delValMap := make(map[string]bool)
		for _, key := range delValList {
			delValMap[key] = true
		}
		newValList := make([]string, 0, len(oldValList))
		for _, val := range oldValList {
			if !delValMap[val] {
				newValList = append(newValList, val)
			}
		}
		newVal = strings.Join(newValList, splitKey)
	}
	return newVal
}
