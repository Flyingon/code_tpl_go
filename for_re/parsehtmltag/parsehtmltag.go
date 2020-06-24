package main

import (
	"fmt"
	"regexp"
)

func GetTagData(data, tagStr string) {
	regStr := fmt.Sprintf("<%s>(.*?)</%s>", tagStr, tagStr)
	fmt.Println("reg: ", regStr)
	re, _ := regexp.Compile(regStr)
	out := re.FindAllStringSubmatch(data, -1)

	for _, i := range out {
		fmt.Println(i[1])
	}
}

func main() {
	GetTagData("书籍UIC<em>测</em><em>试</em>", "em")
}