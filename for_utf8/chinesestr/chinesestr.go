package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "ABC我是＂＃＄中国人）＊＋，"
	size := utf8.RuneCountInString(text)
	fmt.Println("utf8 len:", size)
	strRune := []rune(text)
	fmt.Println(string(strRune[3]) == "我")
	fmt.Println(string(strRune))
	for i := 0; i < len(strRune); i++ {
		fmt.Printf("strRune[%d]=%v\n", i, string(strRune[i]))
	}
	for i, ch := range []byte(text) {
		fmt.Printf("%.2d：%s\n", i, string(ch)) //ch为代表Unicode字符的rune类型
		if i == 0 {
			fmt.Println(string(ch) == "我")
		}
		if i == 27 {
			fmt.Println(string(ch) == "＊")
		}
	}
}
