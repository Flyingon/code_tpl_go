package main

import (
	"fmt"
	"sync"
	"time"
)

var A = sync.Mutex{}
var B = sync.Mutex{}

// goroutineA 协程函数 A
func goroutineA() {
	fmt.Printf("thread A waiting get ResourceA \n")
	A.Lock()
	fmt.Printf("thread A got ResourceA \n")

	time.Sleep(1 * time.Second)

	fmt.Printf("thread A waiting get ResourceB \n")
	B.Lock()
	fmt.Printf("thread A got ResourceB \n")

	A.Unlock()
	B.Unlock()
	return
}

// goroutineB 协程函数 B
func goroutineB() {
	fmt.Printf("thread B waiting get ResourceB \n")
	B.Lock()
	fmt.Printf("thread B got ResourceB \n")

	time.Sleep(1 * time.Second)

	fmt.Printf("thread B waiting get ResourceA \n")
	A.Lock()
	fmt.Printf("thread B got ResourceA \n")

	A.Unlock()
	B.Unlock()
	return
}

// goroutineBSolved 协程函数 B
func goroutineBSolved() {
	fmt.Printf("thread B waiting get ResourceA \n")
	A.Lock()
	fmt.Printf("thread B got ResourceA \n")

	time.Sleep(1 * time.Second)

	fmt.Printf("thread B waiting get ResourceB \n")
	B.Lock()
	fmt.Printf("thread B got ResourceB \n")

	A.Unlock()
	B.Unlock()
	return
}

func main() {
	//创建两个协程
	go func() {
		goroutineA()
		fmt.Println("A finish")
	}()
	//go func() {
	//	goroutineB()
	//	fmt.Println("B finish")
	//}()
	go func() {
		goroutineBSolved()
		fmt.Println("B finish")
	}()

	<-time.After(3 * 60 * time.Second)
	return
}
