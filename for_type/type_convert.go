package main

import (
	"../util"
	"fmt"
)

func main() {
	var oriData interface{}
	data := util.GetStringFromInterface(oriData)
	fmt.Println(data)
}