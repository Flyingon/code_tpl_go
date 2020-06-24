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

// StrToTime 字符串解析到时间
func StrToTime (msgTimeStr string)  {
	msgTime, err := time.ParseInLocation("200601021504", msgTimeStr, time.Local)
	if err != nil {
		fmt.Printf("[ERROR] parse msgTime err: %v, msgTime: %s", err, msgTimeStr)
		return
	}
	fmt.Printf("time: %v", msgTime)
}


func main() {
	timeCost()
	fmt.Printf("1578041406 to now pass: %0.2f ms\n", TsToNow(1578041406))
	StrToTime("201911281609")

	ts1 := int64(1589949475)
	ts2 := time.Now().Unix()
	fmt.Printf("%d, %d, %d\n", ts1, ts2, ts2-ts1)
}
