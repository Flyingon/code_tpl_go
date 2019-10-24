package main

import "fmt"

func main() {
	c := []string{
		"a", "b", "c",
	}
	a := make([]string, len(c))
	copy(a, c)
	fmt.Println(a)
}
