package main

import (
	"fmt"
	"strconv"
	"time"
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

// 获取time当天时间零点
func GetTimeBegin(t time.Time) time.Time {
	startTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return startTime
}

// 获取ts当天时间23:59:59
func GetTimeEnd(t time.Time) time.Time {
	endTime := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return endTime
}

// GetBeforeTs 获取n天前的时间戳
func GetBeforeTs(timeStr string) (tsStart, tsEnd time.Time, err error) {
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
	tsStart = GetTimeBegin(getTime)
	tsEnd = GetTimeEnd(nowTime)
	return
}

// GetAfterTimeMaxHour
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func GetAfterTimeMaxHour(timeStr string) (time.Time, error) {
	d, err := time.ParseDuration(timeStr)
	nowTime := time.Now()
	if err != nil {
		return nowTime, err
	}
	return nowTime.Add(d), nil
}

// GetBeforeTs 获取n天后的时间
func GetAfterTime(timeStr string) (time.Time, error) {
	nowTime := time.Now()
	numStr := timeStr[:len(timeStr)-1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return nowTime, err
	}
	unit := timeStr[len(timeStr)-1:]
	getTime := nowTime
	switch unit {
	case "y":
		getTime = nowTime.AddDate(+num, 0, 0)
	case "m":
		getTime = nowTime.AddDate(0, +num, 0)
	case "d":
		getTime = nowTime.AddDate(0, 0, +num)
	default:
		err = fmt.Errorf("unspport time format: %s", timeStr)
	}
	return getTime, nil
}

// GetCurrentTimeKey 每天24小时根据interval分片，获取当前时间所在片的key，interval最好可以被3整除
func GetCurrentTimeKey(interval int) string {
	curTime := time.Now()
	deltaHour := curTime.Hour()

	if interval == 0 {
		interval = 1
	}
	passUnits := deltaHour / interval
	retTime := time.Date(curTime.Year(), curTime.Month(), curTime.Day(), passUnits*interval, 0, 0, 0, curTime.Location())

	return retTime.Format("2006010215")
}

// GetBeforeTime 获取n天前的时间
func GetBeforeTime(deltaStr string) (timeStr time.Time, err error) {
	nowTime := time.Now()
	numStr := deltaStr[:len(deltaStr)-1]
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return
	}
	unit := deltaStr[len(deltaStr)-1:]
	getTime := nowTime
	switch unit {
	case "y":
		getTime = nowTime.AddDate(-num, 0, 0)
	case "m":
		getTime = nowTime.AddDate(0, -num, 0)
	case "d":
		getTime = nowTime.AddDate(0, 0, -num)
	case "h":
		getTime = nowTime.Add(-time.Duration(num) * time.Hour)
	default:
		err = fmt.Errorf("unspport input format: %s", deltaStr)
	}
	return GetTimeBegin(getTime), err
}

// checkNowInDuration 检查当前时间是否在时间范围内: startTime-endTime
func checkNowInRange(startTime, endTime, fmtStr string) (bool, error) {
	sTime, err := time.ParseInLocation(fmtStr, startTime, time.Local)
	if err != nil {
		return false, err
	}
	eTime, err := time.ParseInLocation(fmtStr, endTime, time.Local)
	if err != nil {
		return false, err
	}
	nTime := time.Now()
	if nTime.Sub(sTime) < 0 || nTime.Sub(eTime) > 0 {
		return false, nil
	}
	return true, nil
}

// GetDeltaDayTs 获取时间戳curTs的n天差时间前后的时间戳
func GetDeltaDayTs(curTs int64, delta int) (tsStart, tsEnd int64, err error) {
	curTime := time.Unix(curTs, 0)
	getTime := curTime.AddDate(0, 0, delta)
	NextTime := curTime.AddDate(0, 0, delta+1)
	tsStart = GetTimeBegin(getTime).Unix()
	tsEnd = GetTimeBegin(NextTime).Unix()
	return
}

// GetDeltaWeekZeroTs 获取n周前0点时间戳
func GetDeltaWeekZeroTs(delta int) (ret int64) {
	now := time.Now()

	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	offset = offset - delta*7
	fmt.Println("offset: ", offset)
	weekStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).Unix()
	return weekStart
}

func main() {
	//fmt.Println("today begin: ", GetTodayBegin())
	//fmt.Println("ts time: ", GetTimeFormat(1558430118))
	//fmt.Println("ts begin: ", GetTsBegin(1558430118))
	//fmt.Println(GetBeforeTs("0d"))
	//
	//timeStamp := 1578041406
	//timeNow := time.Now()
	//timeTs := time.Unix(int64(timeStamp), 0)
	//timeSub := timeNow.Sub(timeTs)
	//timeSubSec := timeSub.Seconds()
	//fmt.Println(timeSubSec)
	//
	//fmt.Printf("--------------%s-------------\n", "GetAflterTs")
	//fmt.Println(GetAfterTimeMaxHour("6Month"))
	//fmt.Printf("--------------%s-------------\n", "GetAfterTime")
	//fmt.Println(GetAfterTime("3d"))
	//fmt.Printf("--------------%s-------------\n", "GetCurrentTimeKey")
	//fmt.Println(GetCurrentTimeKey(4))
	//
	//fmt.Printf("--------------%s-------------\n", "GetBeforeTimeStr")
	//t, e := GetBeforeTime("24h")
	//fmt.Println(t.Unix(), e)

	//ret, e := checkNowInRange("2020-07-01T16:00:00.000Z", "2020-07-01T16:00:00.000Z", "2006-01-02T15:04:05.000Z")
	//fmt.Println(ret, e)
	//begin, end, _ := GetDeltaDayTs(time.Now().Unix(), -1)
	//fmt.Println((end - 1 - begin) / 86400)
	//now := time.Now()
	//offset := int(time.Monday - now.Weekday())
	//if offset > 0 {
	//	offset = -6
	//}
	//weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	//for t:=weekStartDate; now.Sub(t) > 0; t = t.Add(24 * time.Hour) {
	//	fmt.Println(t.Format("20060102"))
	//}
	//for t := time.Monday; time.Now().Sub(t) < 0; t
	fmt.Println(GetDeltaWeekZeroTs(1))
}
