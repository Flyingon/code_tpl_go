package main

import (
	"fmt"
)

func arrayInit () {
	var array []string
	fmt.Println(len(array))
}

// 数组拷贝
func arrayCopy() {
	array1 := []string{"1", "2", "3"}
	fmt.Printf("array1 addr: %p\n", &array1)
	array2 := array1
	fmt.Printf("array2 addr: %p\n", &array2)
	array2 = append(array2, "4")
	fmt.Printf("array1: %v, array2: %v\n", array1, array2)
}

func addElem(array *[]string) {
	arrayTemp := *array
	arrayTemp = append(arrayTemp, "a", "b")
}

func main() {
	//arrayInit()
	//intList := make([]int, 0, 10)
	//intList = append(intList, 1)
	//intList = append(intList, 2)
	//fmt.Println(len(intList), intList)
	var list2 []string
	fmt.Println(len(list2))
	addElem(&list2)
	fmt.Println(list2)
	//arrayCopy()
}
