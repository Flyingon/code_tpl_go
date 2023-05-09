package main

import (
	"fmt"
	"time"
)

// HowLong 验证可以持续多久的结构体，所以叫 how long
type HowLong struct {
	A int
	B string
}

// ReturnHowLong 创建一个 HowLong 的指针
func ReturnHowLong() (string, *HowLong) {
	return "new", &HowLong{
		A: 1,
		B: "one",
	}
}

// PrintHowLong 打印
func PrintHowLong(prefix, text string, long *HowLong) {
	if long == nil {
		fmt.Printf("[%s] text: %s, howLong is nil, address.text: %p, address.long: %p\n", prefix, text, &text, long)
	} else {
		fmt.Printf("[%s] text: %s, %+v, address.text: %p, address.long: %p\n", prefix, text, *long, &text, long)
	}
}

func CheckDefer() {
	var text string
	var long *HowLong
	defer func() {
		PrintHowLong("func defer", text, long)
		go func() {
			time.Sleep(1 * time.Second)
			PrintHowLong("func defer async", text, long)
		}()
	}()
	defer PrintHowLong("just defer ", text, long)
	text, long = ReturnHowLong()
	PrintHowLong("no defer", text, long)
}

func main() {
	CheckDefer()
	<-time.After(3 * time.Second)
}
