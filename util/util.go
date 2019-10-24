package util

import (
	"bufio"
	"os"
)

func NewInt32(x int32) *int32 {
	return &x
}

func NewUint32(x uint32) *uint32 {
	return &x
}

func NewString(s string) *string {
	return &s
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

func ReadLines(path string) ([]string, error) {
	inFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer inFile.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
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
