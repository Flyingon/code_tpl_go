package main

import (
	"fmt"
	"time"
)

type BBBB struct {
	BBB string
}

type AAAA struct {
	AAA     string
	PtrBBBB *BBBB
}

// 在外面修改协程中使用的结构体，协程中结构体值也相应变化
func main() {
	a := &AAAA{"AAA", &BBBB{"BBBB"}}
	fmt.Printf("%p %s\n", &a, a.PtrBBBB.BBB)
	go func(a *AAAA) {
		time.Sleep(1 * time.Second)
		fmt.Printf("%p %s", &a, a.PtrBBBB.BBB)
	}(a)
	a.PtrBBBB = &BBBB{BBB: "DDDDD"}
	//a = &AAAA{"BBB", &BBBB{"CCCC"}}
	//fmt.Printf("%p %s", a, a.PtrBBBB.BBB)
	time.Sleep(2 * time.Second)
}
