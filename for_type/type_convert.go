package main

import (
	"fmt"
	"github.com/Flyingon/go-common/util"
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
