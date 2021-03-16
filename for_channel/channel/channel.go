package main

import (
	"fmt"
	"time"
)

var testChan = make(chan string, 1)

func init() {
	go func() {
		for {
			select {
			case channel := <-testChan:
				fmt.Printf("cron update, channel: %s", channel)
				//default:
				//	fmt.Printf("AAAAAAA")
			}
			time.Sleep(3 * time.Second)

		}
	}()
}

func main() {
	for {
		testChan <- time.Now().String()
		time.Sleep(1 * time.Second)
	}
}
