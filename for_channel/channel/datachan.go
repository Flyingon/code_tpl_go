package main

import (
	"fmt"
	"time"
)

var dataChan = make(chan string, 100)

func init() {
	go func() {
		for {
			select {
			case data, ok := <-dataChan:
				fmt.Printf("data: %s, ok: %v\n", data, ok)
			default:
				fmt.Printf("AAAAAAA\n")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
}

func main() {
	go func() {
		dataChan <- time.Now().String()
		fmt.Printf("send1\n")
		time.Sleep(1 * time.Second)
		dataChan <- time.Now().String()
		fmt.Printf("send2\n")
		time.Sleep(1 * time.Second)
		dataChan <- time.Now().String()
		fmt.Printf("send3\n")
		time.Sleep(1 * time.Second)
		close(dataChan)
	}()
	<-time.After(10 * time.Second)
}
