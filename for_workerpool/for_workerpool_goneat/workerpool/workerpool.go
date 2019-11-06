package workerpool

import (
	"runtime"
	"sync"
	"time"
)

// WorkerPool defines the workerpool interface.
//
// For incoming requests, If we create a new goroutine to handle each
// request, when finished goroutine destroies, this manner cannot scale
// well, if number of req is large, number of goroutines will be huge,
// cost of the userspace context switch can be a burden, also memory
// space is limited.
// But if we use workerpool, we can control the number of goroutines
// to satisfy our requirements both of performance and cost.
type WorkerPool interface {
	Start()
	Submit(func()) bool
	Stop()
}

// nopoolWorkerPool is a simple implemention, in fact there's no pool mechanism in it.
type nopoolWorkerPool struct {
}

func (p *nopoolWorkerPool) Start() {}
func (p *nopoolWorkerPool) Submit(f func()) bool {
	f()
	return true
}
func (p *nopoolWorkerPool) Stop() {}

// NewSimplePool create a simple (nopoolWorkerPool) workerpool.
func NewSimplePool() WorkerPool {
	return &nopoolWorkerPool{}
}

// Handler defines the handler (term `task` maybe clearer) that workerpool handles
type Handler = func()

// NewFILOWorkerPool create a new FILO workerpool, last goroutine we used to serve
// the Handler (run the task) maybe used more likely than other goroutines in pool.
func NewFILOWorkerPool(maxWorkersCount int, maxIdleWorkerDuration time.Duration) WorkerPool {
	return &FILOWorkerPool{
		MaxWorkersCount:       maxWorkersCount,
		MaxIdleWorkerDuration: maxIdleWorkerDuration,
	}
}

// FILOWorkerPool workerPool serves incoming connections via a pool of workers
// in FILO order, i.e. the most recently stopped worker will serve the next
// incoming connection, such a scheme keeps CPU caches hot (in theory).
type FILOWorkerPool struct {
	MaxWorkersCount int

	LogAllErrors bool

	MaxIdleWorkerDuration time.Duration

	lock         sync.Mutex
	workersCount int
	mustStop     bool

	ready []*workerChan

	stopCh chan struct{}

	workerChanPool sync.Pool
}

type workerChan struct {
	lastUseTime time.Time
	ch          chan Handler
}

// Start starts the workerpool, before calling wp.Submit(h Handle), wp.Start() should called.
func (wp *FILOWorkerPool) Start() {
	if wp.stopCh != nil {
		panic("BUG: workerPool already started")
	}
	wp.stopCh = make(chan struct{})
	stopCh := wp.stopCh
	go func() {
		var scratch []*workerChan
		for {
			wp.clean(&scratch)
			select {
			case <-stopCh:
				return
			default:
				time.Sleep(wp.getMaxIdleWorkerDuration())
			}
		}
	}()
}

// Stop stops the workerpool
func (wp *FILOWorkerPool) Stop() {
	if wp.stopCh == nil {
		panic("BUG: workerPool wasn't started")
	}
	close(wp.stopCh)
	wp.stopCh = nil

	// Stop all the workers waiting for incoming connections.
	// Do not wait for busy workers - they will stop after
	// serving the connection and noticing wp.mustStop = true.
	wp.lock.Lock()
	ready := wp.ready
	for i, wc := range ready {
		wc.ch <- nil
		ready[i] = nil
	}
	wp.ready = ready[:0]
	wp.mustStop = true
	wp.lock.Unlock()
}

func (wp *FILOWorkerPool) getMaxIdleWorkerDuration() time.Duration {
	if wp.MaxIdleWorkerDuration <= 0 {
		return 10 * time.Second
	}
	return wp.MaxIdleWorkerDuration
}

// clean remove all of the idle workerChan.
func (wp *FILOWorkerPool) clean(scratch *[]*workerChan) {
	maxIdleWorkerDuration := wp.getMaxIdleWorkerDuration()

	// Clean least recently used workers if they didn't serve connections
	// for more than maxIdleWorkerDuration.
	currentTime := time.Now()

	wp.lock.Lock()
	ready := wp.ready
	n := len(ready)
	i := 0
	for i < n && currentTime.Sub(ready[i].lastUseTime) > maxIdleWorkerDuration {
		i++
	}
	*scratch = append((*scratch)[:0], ready[:i]...)
	if i > 0 {
		m := copy(ready, ready[i:])
		for i = m; i < n; i++ {
			ready[i] = nil
		}
		wp.ready = ready[:m]
	}
	wp.lock.Unlock()

	// Notify obsolete workers to stop.
	// This notification must be outside the wp.lock, since ch.ch
	// may be blocking and may consume a lot of time if many workers
	// are located on non-local CPUs.
	tmp := *scratch
	for i, wc := range tmp {
		wc.ch <- nil
		tmp[i] = nil
	}
}

// Submit submit a new Handler (task) to the FILOWorkerPool.
//
// In one FILOWorkerPool, there're many workerCh, workerCh.ch holds the Handler (tasks).
// Each workerCh is served by a goroutine, we control the maximum number of goroutines
// to reduce userspace context switch cost and memory taken up.
// We can only submit task to `ready` state workerCh, when we say `ready`, it means the
// workerCh is put into workerChanPool already and not idle.
func (wp *FILOWorkerPool) Submit(h Handler) bool {
	wc := wp.getWorkerChan()
	if wc == nil {
		return false
	}
	wc.ch <- h
	return true
}

var workerChanCap = func() int {
	// Use blocking workerChan if GOMAXPROCS=1.
	// This immediately switches Serve to WorkerFunc, which results
	// in higher performance (under go1.5 at least).
	if runtime.GOMAXPROCS(0) == 1 {
		return 0
	}

	// Use non-blocking workerChan if GOMAXPROCS>1,
	// since otherwise the Serve caller (Acceptor) may lag accepting
	// new connections if WorkerFunc is CPU-bound.
	return 1
}()

// getWorkerChan return a workerChan from pool or create a new one,
// when we create a new workerChan, we start a goroutine to traverse
// all Handlers (tasks) submitted to it and execute them.
func (wp *FILOWorkerPool) getWorkerChan() *workerChan {
	var wc *workerChan
	createWorker := false

	wp.lock.Lock()
	ready := wp.ready
	n := len(ready) - 1
	if n < 0 {
		if wp.workersCount < wp.MaxWorkersCount {
			createWorker = true
			wp.workersCount++
		}
	} else {
		wc = ready[n]
		ready[n] = nil
		wp.ready = ready[:n]
	}
	wp.lock.Unlock()

	if wc == nil {
		if !createWorker {
			return nil
		}
		obj := wp.workerChanPool.Get()
		if obj == nil {
			obj = &workerChan{
				ch: make(chan Handler, workerChanCap),
			}
		}
		wc = obj.(*workerChan)
		go func() {
			wp.workerFunc(wc)
			wp.workerChanPool.Put(obj)
		}()
	}
	return wc
}

// release try to put this workerChan into ready state in order to receive new Handle (task).
func (wp *FILOWorkerPool) release(wc *workerChan) bool {
	wc.lastUseTime = time.Now().Truncate(time.Second)
	wp.lock.Lock()
	if wp.mustStop {
		wp.lock.Unlock()
		return false
	}
	wp.ready = append(wp.ready, wc)
	wp.lock.Unlock()
	return true
}

// workerFunc traverse Handlers (tasks) submitted to workerChan.ch to execute them.
func (wp *FILOWorkerPool) workerFunc(wc *workerChan) {
	var h Handler

	for h = range wc.ch {
		if h == nil {
			break
		}
		h()

		if !wp.release(wc) {
			break
		}
	}

	wp.lock.Lock()
	wp.workersCount--
	wp.lock.Unlock()
}

// BaseWorkerPool is the base implementation of WorkerPool
type BaseWorkerPool struct {
	workerCount int
	maxSize     int
	sends       chan Handler

	// Each active sender takes out a read lock on this, and when we want to destroy this bundle and clean its
	// workers, we take a write lock
	deletionLock *sync.RWMutex

	disposed     chan bool
	creationTime time.Time
}

// NewWorkerPool builds a new BaseWorkerPool and return it as a WorkerPool. This is the default pool factory.
func NewWorkerPool(maxSize int) WorkerPool {
	return &BaseWorkerPool{
		sends:        make(chan Handler, 100),
		maxSize:      maxSize,
		deletionLock: &sync.RWMutex{},
		disposed:     make(chan bool),
		workerCount:  0,
		creationTime: time.Now(),
	}
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// Submit an item of Work to be executed.
func (p *BaseWorkerPool) Submit(w Handler) bool {
	p.sends <- w
	return true
}

// Start start the workerpool, it's not thread-safe, lock above this
func (p *BaseWorkerPool) Start() {
	// Spawn as many new workers as there are messages in this send, up until workerPoolMaxSize total
	// spawned workers. This way, when there are clients that are only ever triggering a single push at a time,
	// we only ever spawn a single worker, but when there are clients sending a blast to their entire userbase at
	// once, we'll spawn workerPoolMaxSize workers.
	newWorkers := p.maxSize - p.workerCount
	if newWorkers > 0 {
		p.workerCount += newWorkers
		// Build a fixed-size sender pool for this bundle. Each worker in the sender pool loops indefinitely,
		// processing all the sends for this client, effectively throttling the number of simultaneous sends for a given
		// client.
		for i := 0; i < newWorkers; i++ {
			go func() {
				for {
					select {
					case send := <-p.sends:
						send()
					case <-p.disposed:
						return
					}
				}
			}()
		}
	}
}

func (p *BaseWorkerPool) reserve() bool {
	p.deletionLock.RLock()
	select {
	case <-p.disposed:
		p.deletionLock.RUnlock()
		return false
	default:
		return true
	}
}

func (p *BaseWorkerPool) release() {
	p.deletionLock.RUnlock()
}

func (p *BaseWorkerPool) age() time.Duration {
	return time.Since(p.creationTime)
}

// Stop shutdown and dispose this pool, stop the workers and release any shared resources.
func (p *BaseWorkerPool) Stop() {
	p.deletionLock.Lock()
	defer p.deletionLock.Unlock()

	// On the one hand, I don't expect dispose to be called on a single bundle multiple times.
	// On the other hand, it's so easy and safe to defend against that I might as well.
	select {
	case <-p.disposed:
		return
	default:
		close(p.disposed)
	}
}
