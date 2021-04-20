// Package workerpool 协程池
package workerpool

import (
	"errors"
	"sync"
)

// ErrWorkerPoolFull 请求队列满时，丢弃请求
var ErrWorkerPoolFull = errors.New("Job queue is full now, this Job will be droped")

// Job 执行具体任务封装
type Job func()

// New 根据选项新建一个协程池
func New(opts ...Option) *WorkerPool {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}
	jobQueue := make(chan Job, o.maxQueueSize)
	workerQueue := make(chan *worker, o.maxWorkersCount)
	w := &WorkerPool{
		opts:        o,
		jobQueue:    jobQueue,
		workerQueue: workerQueue,
		stop:        make(chan struct{}),
		wg:          &sync.WaitGroup{},
	}
	w.start()
	return w
}

// start 启动协程池
func (wp *WorkerPool) start() {
	for i := 0; i < cap(wp.workerQueue); i++ {
		worker := wp.newWorker()
		worker.start()
	}
	go wp.dispatch()
}

// WorkerPool 协程池实现
type WorkerPool struct {
	opts        *options
	jobQueue    chan Job
	workerQueue chan *worker
	wokers      []*worker
	stop        chan struct{}
	wg          *sync.WaitGroup
}

// worker 执行真正代码逻辑的载体
type worker struct {
	workerQueue chan *worker
	jobChannel  chan Job
	stop        chan struct{}
	wg          *sync.WaitGroup
}

// Start 启动worker
func (w *worker) start() {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		var job Job
		for {
			w.workerQueue <- w
			select {
			case job = <-w.jobChannel:
				job()
			case <-w.stop:
				return
			}
		}
	}()
}

func defaultOptions() *options {
	return &options{
		maxWorkersCount: 1000,
		maxQueueSize:    1000000,
		drop:            true,
	}
}

// options workerpool选项
type options struct {
	maxWorkersCount int
	maxQueueSize    int
	drop            bool
}

// Option 可选参数
type Option func(*options)

// WithMaxWorkersCount 设置最大工作协程数
func WithMaxWorkersCount(num int) Option {
	return func(o *options) {
		o.maxWorkersCount = num
	}
}

// WithMaxQueueSize 设置队列最大数量
func WithMaxQueueSize(num int) Option {
	return func(o *options) {
		o.maxQueueSize = num
	}
}

// WithDropQueue 设置是否队列满丢弃
func WithDropQueue(drop bool) Option {
	return func(o *options) {
		o.drop = drop
	}
}

// Run 通过协程池执行待处理任务
// 如果工作池满，返回ErrWorkerPoolFull错误，f将被丢弃，业务方可自行判断是否进行等待重试
func (wp *WorkerPool) Run(job Job) error {
	if wp.opts.drop {
		select {
		case wp.jobQueue <- job:
			return nil
		default:
			return ErrWorkerPoolFull
		}
	}
	wp.jobQueue <- job
	return nil
}

// newWorker 新建一个woker
func (wp *WorkerPool) newWorker() *worker {
	return &worker{
		workerQueue: wp.workerQueue,
		jobChannel:  make(chan Job),
		stop:        wp.stop,
		wg:          wp.wg,
	}
}

// put 向woker分配一个job
func (w *worker) put(job Job) {
	w.jobChannel <- job
}

// dispatch 分发任务
func (wp *WorkerPool) dispatch() {
	for {
		select {
		case job := <-wp.jobQueue:
			// 取出Worker
			worker := <-wp.workerQueue
			worker.put(job)
		case <-wp.stop:
			return
		}
	}
}

// Release 释放协程池
func (wp *WorkerPool) Release() {
	close(wp.stop)
	wp.wg.Wait()
}
