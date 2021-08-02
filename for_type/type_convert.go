package main

import (
	"code_tpl_go/util"
	"fmt"
	"reflect"
)

func main() {
	var oriData interface{}
	data := util.InterfaceToString(oriData)
	fmt.Println(data, reflect.TypeOf(data))

	for _, testStr := range []string{"1", "2", "0", "1626872874"} {
		fmt.Println(util.InterfaceToInt(testStr))
	}

	for _, testStr := range []float64{2.1, 2.33, 2.3523} {
		fmt.Println(util.InterfaceToString(testStr))
	}
}
