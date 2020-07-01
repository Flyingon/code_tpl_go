package main

import (
	"fmt"
	"time"
)

func d() {
	defer func() {
		fmt.Println("b closed")
	}()
	data := map[string]interface{}{
		"a": "123",
	}
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("data: ", data)
		}
	}()
	return
}

// 协程b中启动协程a，当协程b结束后，协程a可以继续执行，但当主进程结束后，则都结束
// 协程b中启动协程a，如果协程a中发生panic，整个程序panic
func main() {
	go d()
	time.Sleep(100 * time.Second)
}
