package main

import (
	"context"
	"fmt"
	"time"
)
var key string = "Hello word!"

func run(ctx context.Context){
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
		default:
			go JustPrint()
			//fmt.Println(ctx.Value(key))
		}
	}
}

func run2(ctx context.Context) () {
	defer fmt.Println("run2 return")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			JustPrint()
			//time.Sleep(10*time.Second)
			return
		}
	}
}

func JustPrint() {
	i := 0
	for {
		//if i >= 3 {
		//	break
		//}
		fmt.Println(time.Now())
		//time.Sleep(1*time.Second)
		i ++
	}
}

func main() {
	//<-time.After(3 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(),3 * time.Second)
	defer cancel()
	//value := context.WithValue(ctx,key,"This is my test")
	go run(ctx)
	//go run2(ctx)
	time.Sleep(10 * time.Second)
}