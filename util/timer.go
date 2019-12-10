package util

import (
	"time"
	)

// 按月定时
func StartMonthlyTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下个月零点
			year := now.Year()
			month := now.Month()
			if month == time.December {
				year += 1
				month = time.January
			} else {
				month = time.Month(int(month) + 1)
			}
			next := time.Date(year, month, 1, 0, 0, 0, 0, time.Now().Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

// 按分定时
func StartMinuteTimer(intMin uint, f func(interface{}), i interface{}) {
	go func() {
		for {
			f(i)
			now := time.Now()
			next := now.Add(time.Duration(intMin) * time.Minute)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

// 按秒定时
// 每次不重新生成timer
func StartSecondTimer(intSec uint, f func(interface{}), i interface{}) {
	go func() {
		t := time.NewTimer(time.Duration(intSec) * time.Second)
		for {
			f(i)
			<-t.C
			t.Reset(time.Duration(intSec) * time.Second)
		}
	}()
}
