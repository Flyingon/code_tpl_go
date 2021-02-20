package main
import (
	"fmt"
	"sync"
	"time"
)


type SimplePool struct {
	wg   sync.WaitGroup
	work chan func() //任务队列
}

func NewSimplePoll(workers int) *SimplePool {
	p := &SimplePool{
		wg:   sync.WaitGroup{},
		work: make(chan func()),
	}
	p.wg.Add(workers)
	//初始化协程池（添加任务前）时就根据指定的并发量去读取管道并执行，以免添加任务时管道阻塞
	for i := 0; i < workers; i++ {
		go func() {
			defer p.wg.Done()
			defer func() {
				// 捕获异常 防止waitGroup阻塞
				if err := recover(); err != nil {
					fmt.Println(err)
				}
			}()
			// 从workChannel中取出任务执行
			for fn := range p.work {
				fn()
			}
		}()
	}
	return p
}
// 添加任务
func (p *SimplePool) Add(fn func()) {
	p.work <- fn
}

// 执行
func (p *SimplePool) Run() {
	close(p.work)
	p.wg.Wait()
}

// 测试使用
func main() {
	p := NewSimplePoll(20)
	for i := 0; i < 100; i++ {
		p.Add(parseTask(i))
	}
	p.Run()
}

func parseTask(i int) func() {
	return func() {
		// 模拟抓取数据的过程
		time.Sleep(time.Second * 1)
		fmt.Println(time.Now().Unix(), " finish parse ",i)
	}
}