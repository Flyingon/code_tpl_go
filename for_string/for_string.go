package main

import (
	"fmt"
	"strings"
)

func CamelName(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	fmt.Println(name)
	name = strings.Title(name)
	fmt.Println(name)
	return strings.Replace(name, " ", "", -1)
}

func main() {
	pbName := "helloworld_123"
	fmt.Println(CamelName(pbName))
}
