package main

import (
	"fmt"
	"time"
)

func asyncFunc(stopCh chan bool) {
	go func() {
		for {
			select {
			case <-stopCh:
				fmt.Println("stop")
				return
			default:
				fmt.Println("running")
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func main() {
	stopCh := make(chan bool)
	go asyncFunc(stopCh)
	i := 1
	for {
		fmt.Println("main: ", i)
		if i == 2 {
			fmt.Println("ddd: ", len(stopCh))
			stopCh <- true
		}
		time.Sleep(3 * time.Second)
		i++
	}
}
