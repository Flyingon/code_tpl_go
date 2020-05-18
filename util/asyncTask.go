package util

import (
	"fmt"
	"runtime"
	"sync"
)

// GoAndWait 封装更安全的多并发调用, 启动goroutine并等待所有处理流程完成，自动recover
// 返回值error返回的是多并发协程里面第一个返回的不为nil的error，主要用于关键路径判断，当多并发协程里面有一个是关键路径且有失败则返回err，其他非关键路径并发全部返回nil
func GoAndWait(handlers ...func() error) (err error) {

	var wg sync.WaitGroup
	for _, f := range handlers {

		wg.Add(1)
		go func(handler func() error) {

			defer func() {
				if e := recover(); e != nil {
					buf := make([]byte, 1024)
					buf = buf[:runtime.Stack(buf, false)]
					fmt.Printf("[PANIC]%v\n%s\n", e, buf)
				}
				wg.Done()
			}()

			e := handler()
			if err == nil && e != nil {
				err = e
			}
		}(f)
	}

	wg.Wait()

	return
}
