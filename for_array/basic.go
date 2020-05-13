package main

import "fmt"

func arrayInit () {
	var array []string
	fmt.Println(len(array))
}

func main() {
	arrayInit()
	intList := make([]int, 0, 10)
	intList = append(intList, 1)
	intList = append(intList, 2)
	fmt.Println(len(intList))
}
