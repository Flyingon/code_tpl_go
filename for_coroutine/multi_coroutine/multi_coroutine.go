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

// 协程b中启动协程a，当协程b结束后，协程a可以继续执行，但当主进程结束后，则都结束
// 协程b中启动协程a，如果协程a中发生panic，整个程序panic
func main() {
	go b()
	time.Sleep(100 * time.Second)
}
