package main

import (
	"fmt"
	"os"
	"os/signal"
)

func listenSignal(signals <-chan os.Signal) {
	for {
		select {
		case d := <-signals:
			fmt.Printf("get signal: %+v\n", d.String())
			fmt.Printf("!!!!!!!!!!")
		}
	}
}

func signal1() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	listenSignal(signals)
}

func signal2() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	listenSignal(signals)
}

func signal3() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	listenSignal(signals)
}

func main() {
	go signal1()
	go signal3()
	signal2()
}
