package workerpool

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewPool测试创建和跑协程池
func TestNewPool(t *testing.T) {
	pool := New(WithMaxWorkersCount(10000))
	defer pool.Release()

	iterations := 1000000
	var counter uint64 = 0

	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		arg := uint64(1)
		job := func() {
			defer wg.Done()
			atomic.AddUint64(&counter, arg)
		}

		pool.Run(job)
	}
	wg.Wait()

	counterFinal := atomic.LoadUint64(&counter)
	assert.Equal(t, uint64(iterations), counterFinal)
}

// TestOption 测试创建option
func TestOption(t *testing.T) {
	opts := &options{}
	WithMaxQueueSize(20000)(opts)
	WithMaxWorkersCount(10000)(opts)
	WithDropQueue(true)(opts)
	assert.Equal(t, opts.drop, true)
	assert.Equal(t, opts.maxQueueSize, 20000)
	assert.Equal(t, opts.maxWorkersCount, 10000)
}

// TestDropQueue 测试队列满丢弃的情况
func TestDropQueue(t *testing.T) {

	// 测试pool配置了WithDropQueue，但队列满的情况
	pool := New(WithMaxWorkersCount(1), WithMaxQueueSize(0), WithDropQueue(true))
	err := pool.Run(func() {})
	assert.Equal(t, err, ErrWorkerPoolFull)

	// 测试pool配置了WithDropQueue，但队列未满的情况
	pool = New(WithMaxWorkersCount(1000), WithMaxQueueSize(1001), WithDropQueue(true))
	defer pool.Release()

	iterations := 1000
	var counter uint64 = 0

	wg := sync.WaitGroup{}
	wg.Add(iterations)
	for i := 0; i < iterations; i++ {
		arg := uint64(1)
		job := func() {
			defer wg.Done()
			atomic.AddUint64(&counter, arg)
		}

		pool.Run(job)
	}
	wg.Wait()

	counterFinal := atomic.LoadUint64(&counter)
	assert.Equal(t, uint64(iterations), counterFinal)
}
