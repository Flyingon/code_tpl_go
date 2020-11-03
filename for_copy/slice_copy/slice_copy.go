package main

import "fmt"

func arrayCopy() {
	array1 := []string{
		"a", "b", "c",
	}
	array2 := array1
	array2 = append(array2, "d")
	fmt.Printf("array1: %+v, addr: %p\n", array1, array1)
	fmt.Printf("array2: %+v, addr: %p\n", array2, array2)
}

func arrayUseCopy() {
	c := []string{
		"a", "b", "c",
	}
	a := make([]string, len(c))
	copy(a, c)
	fmt.Println(a)
}

func mapCopy() {
	map1 := map[string]string {
		"1": "a",
		"2": "b",
		"3": "c",
	}
	map2 := map1
	map2["4"] = "d"
	fmt.Printf("map1: %+v, addr: %p\n", map1, map1)
	fmt.Printf("map2: %+v, addr: %p\n", map2, map2)
}

func main() {
	arrayCopy()
	mapCopy()
}
