package main

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

var rateLimiter = rate.NewLimiter(10, 10)

func PrintNum(num int) {
	fmt.Printf("%d, excute: %d\n", time.Now().Unix(), num)
	time.Sleep(200 * time.Millisecond)
}

func rateLimitTest() {
	total := 50
	wg := sync.WaitGroup{}
	wg.Add(total)
	ctxLimiter := context.Background()
	sTime := time.Now()
	for i := 0; i < total; i++ {
		num := i
		rateLimiter.Wait(ctxLimiter)
		go func() {
			defer wg.Done()
			PrintNum(num)
		}()
	}
	wg.Wait()
	costTime := float32(time.Now().Sub(sTime).Nanoseconds() / 1e6)
	fmt.Println("cost time: ", costTime)
}

func rateLimitCompareTest() {
	total := 50
	wg := sync.WaitGroup{}
	wg.Add(total)
	sTime := time.Now()
	for i := 0; i < total; i++ {
		num := i
		PrintNum(num)
	}
	costTime := float32(time.Now().Sub(sTime).Nanoseconds() / 1e6)
	fmt.Println("cost time: ", costTime)
}

func main() {
	rateLimitTest()
	//rateLimitCompareTest()
}
