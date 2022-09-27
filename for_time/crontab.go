package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

var timeLayout = "2006-01-02_15:04:05"

// getCronNextTimeShow 查询cron下次执行时间
func getCronNextTimeShow(cronSpec string) (string, error) {
	nextTime, err := getCronNextTime(cronSpec)
	if err != nil {
		return "", err
	}
	return nextTime.Format(timeLayout), nil
}

// getCronNextTime 查询cron下次执行时间
func getCronNextTime(cronSpec string) (time.Time, error) {
	tz := "Asia/Hong_Kong"
	loc, _ := time.LoadLocation(tz)
	sch, err := cron.ParseStandard(cronSpec)

	if err != nil {
		return time.Time{}, err
	}
	return sch.Next(time.Now().In(loc)), nil
}

func main() {
	t, err := getCronNextTime("1 0,8 * * *")
	fmt.Println(t.Unix(), err)
}
