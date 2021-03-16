package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	cronSpec := "0 */1 * * * ?"
	//timeLayout := "2006-01-02_15:04:05"

	sch, _:= cron.Parse(cronSpec)
	fmt.Println(time.Now())
	fmt.Println(sch.Next(time.Now()))
	//_ = c.AddFunc(cronSpec, func() {
	//	fmt.Print("now: ", time.Now().Format(timeLayout))
	//	fmt.Printf("\nc.Entries()[0].Prev:%s\nc.Entries()[0].Next:%s",
	//		c.Entries()[0].Prev.Format(timeLayout), c.Entries()[0].Next.Format(timeLayout))
	//})
	//c.Start()

}
