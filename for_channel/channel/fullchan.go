package main

import (
	"fmt"
	"strconv"
	"time"
)

var fullChan = make(chan string, 10)

func init() {
	fmt.Println("chan len: ", len(fullChan))
	go func() {
		for {
			fmt.Println("chan len: ", len(fullChan))
			select {
			case data, ok := <-fullChan:
				fmt.Printf("get data: %v, ok: %v\n", data, ok)
				time.Sleep(time.Millisecond * 600)
			default:
				fmt.Printf("AAAAAAA\n")
			}
			time.Sleep(300 * time.Millisecond)
		}
	}()
}

func main() {
	for i := 0; i <= 10000; i++ {
		fullChan <- strconv.FormatInt(int64(i), 10)
	}
	<-time.After(10 * time.Second)
}
