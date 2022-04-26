package main

import (
	"fmt"
	"sync"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	//fmt.Println(o.done)
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	once := Once{
		m: sync.Mutex{},
	}
	for i := 0; i < 10000000; i++ {
		go once.Do(func() {
			fmt.Printf("%d%d%d%d%d%d\n", i, i, i, i, i, i)
		})
	}
	<-time.After(3 * time.Second)
}
