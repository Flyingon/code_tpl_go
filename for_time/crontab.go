package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

var timeLayout = "2006-01-02_15:04:05"

func main() {
	//cronSpec := "0 */1 * * * *"
	cronSpec := "*/5 * * * * *"
	sch, _ := cron.Parse(cronSpec)
	fmt.Println(time.Now())
	fmt.Println(sch.Next(time.Now()))

	ct := cron.New()
	_ = ct.AddFunc(cronSpec, func() {
		fmt.Print("now: ", time.Now().Format(timeLayout))
		fmt.Printf("\nc.Entries()[0].Prev:%s\nc.Entries()[0].Next:%s",
			ct.Entries()[0].Prev.Format(timeLayout), ct.Entries()[0].Next.Format(timeLayout))
	})
	ct.Start()
	<-time.After(1000 * time.Second)
}
