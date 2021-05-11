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
			fmt.Printf("get signal: %+v", d.String())
			fmt.Sprintf("!!!!!!!!!!")
		}
	}
}

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill)
	listenSignal(signals)
}
