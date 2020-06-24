// Package workerpool 协程池
package workerpool

import (
	"sync"
	"time"
)

// New 根据选项新建一个协程池
func New(opts ...Option) *WorkerPool {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	w := &WorkerPool{opts: o}
	w.pool.New = func() interface{} {
		return &worker{
			runners: make(chan func(), 1),
		}
	}
	return w
}

// WorkerPool 协程池实现
type WorkerPool struct {
	opts *options

	mu sync.Mutex

	pool sync.Pool

	ready        []*worker
	workersCount int
}

type worker struct {
	lastUseTime time.Time
	runners     chan func()
}

func defaultOptions() *options {
	return &options{maxWorkersCount: 10000}
}

type options struct {
	maxWorkersCount int
}

// Option 可选参数
type Option func(*options)

// WithMaxWorkersCount 设置最大工作协程数
func WithMaxWorkersCount(num int) Option {
	return func(o *options) {
		o.maxWorkersCount = num
	}
}

// Run 通过协程池执行待处理任务
func (wp *WorkerPool) Run(f func()) error {

	var w *worker
	var createWorker bool

	wp.mu.Lock()
	ready := wp.ready
	n := len(ready) - 1
	if n < 0 {
		if wp.workersCount < wp.opts.maxWorkersCount {
			createWorker = true
			wp.workersCount++
		}
	} else {
		w = ready[n]
		ready[n] = nil
		wp.ready = ready[:n]
	}
	wp.mu.Unlock()

	if w == nil {
		if !createWorker {
			return nil
		}
		vw := wp.pool.Get()
		w = vw.(*worker)
		go func() {
			wp.runWorker(w)
			wp.pool.Put(vw)
		}()
	}

	w.runners <- f
	return nil
}

func (wp *WorkerPool) runWorker(w *worker) {
	for f := range w.runners {
		if f == nil {
			break
		}
		f()

		w.lastUseTime = time.Now().Truncate(time.Second)
		wp.mu.Lock()
		wp.ready = append(wp.ready, w)
		wp.mu.Unlock()
	}

	wp.mu.Lock()
	wp.workersCount--
	wp.mu.Unlock()
}
