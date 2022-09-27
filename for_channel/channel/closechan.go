package main

import (
	"fmt"
	"time"
)

var closeChan = make(chan struct{})

func init() {
	go func() {
		for {
			select {
			case <-closeChan:
				fmt.Printf("closed")
				return
			default:
				fmt.Printf("excute: %d\n", time.Now().Unix())
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
}

func main() {
	go func() {
		time.Sleep(2 * time.Second)
		close(closeChan)
	}()
	<-time.After(10 * time.Second)
}
