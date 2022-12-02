package main

import (
	"fmt"
	"sync"
	"time"
)

var syncMap sync.Map

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			for {
				a, ok := syncMap.Load("a")
				fmt.Println(a, ok)
				time.Sleep(1 * time.Second)
			}
		}()
	}
	go func() {
		for {
			//syncMap.Store("a", time.Now().Unix())
			newSyncMap := sync.Map{}
			newSyncMap.Store("a", time.Now().Unix())
			syncMap = newSyncMap
			fmt.Println("syncMap refresh")
			time.Sleep(2 * time.Second)
		}

	}()

	<-time.After(3 * 60 * time.Second)
	return
}
