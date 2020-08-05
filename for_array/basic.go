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
	data := map[string]map[string]interface{}{
		"1": {"a": "a"},
		"2": {"b": "b"},
		"3": {"c": "c"},
	}
	dataNew := make([]map[string]interface{}, 0 ,len(data))
	for k, v :=range data {
		fmt.Printf("%s: %p\n", k, &v)
		v["sub_id"] = k
		dataNew = append(dataNew, v)
	}
	fmt.Printf("dataNew: %v", dataNew)
	//var list2 []string
	//fmt.Println(len(list2))
	//addElem(&list2)
	//fmt.Println(list2)
	//arrayCopy()
}
