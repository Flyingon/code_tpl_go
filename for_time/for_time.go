package main

import (
	"strconv"
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

// 获取time当天时间零点时间戳
func GetTimeBeginTs(t time.Time) int64 {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startTime.Unix()
}

// 获取ts当天时间零点时间戳
func GetTimeEndTs(t time.Time) int64 {
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return endTime.Unix()
}

// GetBeforeTs 获取n天前的时间戳
func GetBeforeTs(timeStr string) (tsStart, tsEnd int64, err error) {
	nowTime := time.Now()
	numStr := timeStr[:len(timeStr)-1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	unit := timeStr[len(timeStr)-1:]
	getTime := nowTime
	switch unit {
	case "y":
		getTime = nowTime.AddDate(-num, 0, 0)
	case "m":
		getTime = nowTime.AddDate(0, -num, 0)
	case "d":
		getTime = nowTime.AddDate(0, 0, -num)
	default:
		err = fmt.Errorf("unspport time format: %s", timeStr)
	}
	tsStart = GetTimeBeginTs(getTime)
	tsEnd = GetTimeEndTs(nowTime)
	return
}


func main() {
	fmt.Println("today begin: ", GetTodayBegin())
	fmt.Println("ts time: ", GetTimeFormat(1558430118))
	fmt.Println("ts begin: ", GetTsBegin(1558430118))
	fmt.Println(GetBeforeTs("0d"))
}
