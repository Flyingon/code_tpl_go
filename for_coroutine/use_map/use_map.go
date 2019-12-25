/**
 * 并发编程，切片的线程安全性问题
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var dict map[int]bool = map[int]bool{}
var wgDict sync.WaitGroup = sync.WaitGroup{}
var muDict sync.Mutex = sync.Mutex{}
var rwDict sync.RWMutex = sync.RWMutex{}

// 线程不安全的方法,直接panic
func addNotSafe(i int) {
	dict[i] = true
	wgDict.Done()
}

// 线程安全的方法，增加了互斥锁
func addSafeMul(i int) {
	muDict.Lock()
	dict[i] = true
	muDict.Unlock()
	wgDict.Done()
}

// 线程安全的方法，增加了读写锁
func addSafeRw(i int) {
	rwDict.Lock()
	dict[i] = true
	rwDict.Unlock()
	wgDict.Done()
}


/*
返回结果: 互斥锁反而快很多，需要进一步分析
不加锁直接失败
dict add num=1000000
list len=1000000, time=6422 ms
new dict with rw lock add num=1000000
new dict with mutex lock len=1000000, time=623 ms
*/
func main() {
	// 并发启动的协程数量
	max := 100000
	dict = map[int]bool{}
	fmt.Printf("dict add num=%d\n", max)
	wgDict.Add(max)
	time1 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go addSafeRw(i)
	}
	wgDict.Wait()
	time2 := time.Now().UnixNano()
	fmt.Printf("list len=%d, time=%d ms\n", len(dict), (time2-time1)/1000000)

	// 覆盖后再执行一次
	dict = map[int]bool{}
	fmt.Printf("new dict with rw lock add num=%d\n", max)
	wgDict.Add(max)
	time3 := time.Now().UnixNano()
	for i := 0; i < max; i++ {
		go addSafeMul(i)
	}
	wgDict.Wait()
	time4 := time.Now().UnixNano()
	fmt.Printf("new dict with mutex lock len=%d, time=%d ms\n", len(dict), (time4-time3)/1000000)
}
