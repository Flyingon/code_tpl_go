/**
 * 并发编程，切片的线程安全性问题
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var list []int = []int{}
var wgList sync.WaitGroup = sync.WaitGroup{}
var muList sync.Mutex = sync.Mutex{}

/*
返回结果:
list add num=10000
list len=10000, time=2 ms
*/
func main() {
	// 并发启动的协程数量
	max := 10000
	fmt.Printf("list add num=%d\n", max)

	wgList.Add(max)
	intChan := make(chan int , max)
	time1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go func(c int){
			intChan <- c
			wgList.Done()
		}(i)
	}
	wgList.Wait()
	close(intChan)
	time2 := time.Now().UnixNano()
	// 注意list里面是地址还是值哦，地址就要新建变量
	for i := range intChan{
		list = append(list, i)
	}
	fmt.Printf("list len=%d, time=%d ms\n", len(list), (time2-time1)/1000000)
	//fmt.Println(list)
}
