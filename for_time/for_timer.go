package main

import (
	"fmt"
	"time"
)

func PrintTime(argc interface{}) {
	fmt.Println("now: ", time.Now())
	<-time.After(3 * time.Second)
	return
}

func useTicker() {
	ticker := time.NewTicker(time.Duration(1000*1000/10) * time.Microsecond)
	stopTK := make(chan bool, 1)
	defer ticker.Stop()
	ticks := 0
	for {
		select {
		case <-ticker.C:
			fmt.Println("t1定时器: ", ticks)
			ticks++
			if ticks == 10 {
				stopTK <- true
			}
		case <-stopTK:
			fmt.Println("定时器关闭: ", ticks)
			return
		}
	}
}

func main() {
	useTicker()
}
