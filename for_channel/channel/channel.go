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
				fmt.Printf("cron update, channel: %s\n", channel)
			default:
				fmt.Printf("AAAAAAA\n")
			}
			time.Sleep(3 * time.Second)
		}
	}()
}

func chanFull() {
	for {
		select {
		case testChan <- time.Now().String():
			fmt.Printf("1111111\n")
		default:
			fmt.Printf("AAAAAAA\n")
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	//testChan <- time.Now().String()

	//chanFull()
	for {
		testChan <- time.Now().String()
		fmt.Printf("send1\n")
		testChan <- time.Now().String()
		fmt.Printf("send2\n")
		testChan <- time.Now().String()
		fmt.Printf("send3\n")
		//time.Sleep(1 * time.Second)
	}
}
