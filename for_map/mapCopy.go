package main

import (
	"fmt"
	"github.com/Flyingon/go-common/util"
)

func MapCopyExample() {
	map1 := map[string]interface{}{
		"a": "a",
		"1": 1,
	}
	map2 := util.MapCopy(map1)
	fmt.Println("map1: ", map1)
	fmt.Println("map2: ", map2)
	map3 := util.MapCopy(nil)
	fmt.Println("map3: ", map3)
}

func main() {
	MapCopyExample()
}
