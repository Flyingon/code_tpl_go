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
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	ticks := 0
	for {
		select {
		case <-ticker.C:
			fmt.Println("t1定时器: ", ticks)
			ticks ++
			if ticks == 10 {
				return
			}
		}
	}
}

func main() {
	useTicker()
}
