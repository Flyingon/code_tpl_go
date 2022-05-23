package main

import "fmt"

var tempStr = "abcABC"

func printIndex() {
	for i := 0; i < len(tempStr); i++ {
		fmt.Printf("%v,", tempStr[i])
	}
	fmt.Printf("\n")
}

// SetChar 修改字符串对应字节
func SetChar(targetStr *string, pos int, newChar string) {
	if len(*targetStr) <= pos {
		for i := pos - len(*targetStr); i >= 0; i-- {
			*targetStr += " "
		}
	}
	temp := []rune(*targetStr)
	temp[pos] = []rune(newChar)[0]
	*targetStr = string(temp)
}

func main() {
	//printIndex()
	//SetChar(&tempStr, 0 , "D")
	//fmt.Println(tempStr)
	a := ""
	SetChar(&a, 0, "X")
	fmt.Println(a)
	fmt.Println(tempStr[1] == 'b')
}
