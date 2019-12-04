package main

import(
	"../util"
	"fmt"
	"time"
)

func PrintTime(argc interface{}) {
	fmt.Println("now: ", time.Now())
	<-time.After(3 * time.Second)
	return
}

func main() {
	util.StartSecondTimer(2, PrintTime, nil)
	<-time.After(60 * time.Second)
}
