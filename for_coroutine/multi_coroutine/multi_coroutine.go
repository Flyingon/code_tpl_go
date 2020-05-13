package main

import (
	"fmt"
	"time"
)

func a() {
	panicNum := 3
	i := 0
	for true {
		if i > panicNum {
			panic("a panic")
		}
		fmt.Printf("a time: %v\n", time.Now())
		time.Sleep(1000 * time.Millisecond)
		i ++
	}
}

func b() {
	defer func() {
		fmt.Println("b closed")
	}()
	go a()
	for i := 1; i < 5; i++ {
		fmt.Printf("b time: %v\n", time.Now())
		time.Sleep(1000 * time.Millisecond)
	}
	return
}

func main() {
	go b()
	time.Sleep(100 * time.Second)
}
