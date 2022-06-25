package main

import (
	"code_tpl_go/for_workerpool/for_workerpool_goneat/workerpool"
	"fmt"
	"time"
)

func PrintNum(num int) {
	fmt.Printf("excute: %d\n", num)
}

func FILOWorkerPoolTest() {
	wp := workerpool.NewFILOWorkerPool(10, 2*time.Second)
	wp.Start()
	for i := 0; i < 100; i++ {
		num := i
		wp.Submit(func() {
			PrintNum(num)
		})
	}
	fmt.Println("workerpool close")
	wp.Stop()
	<-time.After(3 * time.Second)
}

func main() {
	FILOWorkerPoolTest()
}
