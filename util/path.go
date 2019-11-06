package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ParentPath string

func init() {
	if path, err := GetParentDirectory(); err == nil {
		ParentPath = path
	} else {
		panic(fmt.Sprintf("get parent path err: %v", err))
	}
}

// Get current directory of executable file,
// if fails returns empty filepath and error
func GetCurrentDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return dir, nil
}

// returns substring of s, i.e., s[pos:pos+length]
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// Get parental directory of executable file,
// if fails returns empty filepath and error
func GetParentDirectory() (string, error) {
	dirctory, err := GetCurrentDirectory()
	if err != nil {
		return "", err
	}
	parentPath := substr(dirctory, 0, strings.LastIndex(dirctory, string(filepath.Separator)))
	return parentPath, nil
}

// Check whether file is existed,
// if existed returns true, otherwise returns false
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
