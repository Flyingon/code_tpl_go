package main

import (
	"fmt"
	"time"
)

// timeCost 计算耗时，毫秒返回
func timeCost() {
	beginTime := time.Now()
	time.Sleep(555 * time.Millisecond)
	costTime := int(time.Since(beginTime) / time.Millisecond)
	fmt.Printf("cost time: %d ms\n", costTime)
}

// TsToNow 时间戳ts到当前的时间
func TsToNow(ts int64) float64 {
	timeStamp := 1578041406
	timeNow := time.Now()
	timeTs := time.Unix(int64(timeStamp), 0)
	timeSub := timeNow.Sub(timeTs)
	timeSubSec := timeSub.Seconds()
	return timeSubSec
}

func main() {
	timeCost()
	fmt.Printf("1578041406 to now pass: %0.2f ms\n", TsToNow(1578041406))
}
