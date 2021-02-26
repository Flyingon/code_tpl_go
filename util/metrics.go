package util

import (
	"fmt"
	"time"
)

// ReportMonitor 上报监控
// args: 0: 上报累计的数值,默认是1
// toDo: 自定义实现
func ReportMonitor(msg string, args ...float64) {
	fmt.Printf("ReportMonitor: %s\n", msg)
}

// ReportTimeDuration 耗时上报
// toDo: 自定义实现
func ReportTimeDuration(msg string, duration time.Duration) {
	fmt.Printf("ReportTimeDuration: %s, duration: %s\n", msg, duration.String())
}
