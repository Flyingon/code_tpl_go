package main

import (
	"time"
	"fmt"
)

// 返回当前时间格式
func GetDateTime() string {
	return time.Now().Local().Format("2006-01-02 15:04:05")
}

// 格式化传入的时间
func GetTimeFormat(t int64) string {
	return time.Unix(t, 0).Format("2006-01-02 15:04:05")
}

// 获取今天的零点时间
func GetTodayBegin() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t.Unix()
}

// 获取今天的最后结束时间
func GetTodayEnd() int64 {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 23:59:59", time.Local)
	return t.Unix()
}

// 获取ts当天时间零点时间戳
func GetTsBegin(ts int64) int64 {
	timeStr := time.Unix(ts, 0).Format("2006-01-02")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", timeStr+" 00:00:00", time.Local)
	return t.Unix()
}

func main() {
	fmt.Println("today begin: ", GetTodayBegin())
	fmt.Println("ts time: ", GetTimeFormat(1558430118))
	fmt.Println("ts begin: ", GetTsBegin(1558430118))
}
