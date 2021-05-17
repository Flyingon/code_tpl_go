package main

import (
	"fmt"
)

func setFunc(s []int) {
	s[0] = 1
	s = append(s, 2)
}
func setMap(m map[string]interface{}) {
	m["1111"] = 100
}
func main() {
	sli := []int{0}
	fmt.Println(sli)
	setFunc(sli)
	fmt.Println(sli)

	m := map[string]interface{}{
		"a": 10,
	}
	fmt.Println(m)
	setMap(m)
	fmt.Println(m)
}
