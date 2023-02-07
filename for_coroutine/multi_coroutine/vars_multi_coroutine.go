package main

import (
	"fmt"
	"time"
)

type SB struct {
	B string
}

type SA struct {
	A  string
	Sb *SB
}

func coroutine1() {
	defer func() {
		fmt.Println("coroutine1 closed")
	}()
	data1 := map[string]interface{}{
		"a": "123",
	}
	data2 := &SA{
		A:  "aaa",
		Sb: &SB{B: "bbb"},
	}
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("data1: ", data1)
			fmt.Println("data2: ", data2.Sb)
		}
	}()
	go func() {
		time.Sleep(2000 * time.Millisecond)
		data2.Sb = &SB{B: "ccc"}
		time.Sleep(1000 * time.Millisecond)
		data1["b"] = 456
	}()
	return
}

// 协程coroutine1中启动协程a，当协程coroutine1结束后，协程a可以继续执行，但当主进程结束后，则都结束
// 协程coroutine1创建的局部变量 data1 和 data2 被协程a调用，当协程coroutine1结束后，data1 和 data2 不会被回收，可以正常使用
// data1 和 data2 可以被其他携程修改，这是协程a中使用的 data1 和 data2 也会更改
func main() {
	go coroutine1()
	time.Sleep(100 * time.Second)
}

/* 执行结果：
coroutine1 closed
data1:  map[a:123]
data2:  &{bbb}
data1:  map[a:123]
data2:  &{bbb}
data1:  map[a:123]
data2:  &{ccc}
data1:  map[a:123 b:456]
data2:  &{ccc}
data1:  map[a:123 b:456]
data2:  &{ccc}
data1:  map[a:123 b:456]
data2:  &{ccc}
data1:  map[a:123 b:456]
data2:  &{ccc}
data1:  map[a:123 b:456]
data2:  &{ccc}
*/
