package connpool

import (
	"context"
	"errors"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"git.code.oa.com/trpc-go/trpc-go/codec"
	"github.com/stretchr/testify/assert"
)

var (
	network  = "tcp"
	address  = "127.0.0.1:21877"
	address2 = "127.0.0.1:21878"
	ch       = make(chan struct{})
	buffer   = []byte("hello world")
)

func init() {
	go simpleTCPServer(ch)
	<-ch
	go simpleTCPServerClose(ch)
	<-ch
}

func simpleTCPServer(ch chan struct{}) {
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	ch <- struct{}{}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			buffer := make([]byte, 256)
			conn.Read(buffer)
		}()
	}
}

func simpleTCPServerClose(ch chan struct{}) {
	l, err := net.Listen(network, address2)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	ch <- struct{}{}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func() {
			conn.Close()
		}()
	}
}

func newPool(maxIdle, maxActive int) *ConnectionPool {
	pool := &ConnectionPool{
		Dial: func(context.Context) (net.Conn, error) {
			return net.Dial(network, address)
		},
		MaxIdle:   maxIdle,
		MaxActive: maxActive,
	}
	pool.RegisterChecker(time.Millisecond*50, pool.defaultChecker)

	return pool
}

func TestPoolGet(t *testing.T) {
	pool := newPool(2, 10)

	_, err := pool.Get(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, pool.active, 1)
	assert.Equal(t, pool.idle.count, 0)
}

func TestPoolGetOnClose(t *testing.T) {
	pool := newPool(2, 10)
	assert.Nil(t, pool.Close())

	ctx := context.Background()
	_, err := pool.Get(ctx)
	assert.Equal(t, err, ErrPoolclosed)
}

func TestPoolConcurrentGet(t *testing.T) {
	maxActive := 10
	pool := newPool(2, maxActive)
	defer pool.Close()

	var wg sync.WaitGroup
	ctx := context.Background()
	for i := 0; i < maxActive; i++ {
		wg.Add(1)
		func() {
			_, err := pool.Get(ctx)
			assert.Nil(t, err)
			wg.Done()
		}()
	}

	wg.Wait()

	assert.Equal(t, pool.active, maxActive)
	assert.Equal(t, pool.idle.count, 0)
}

func TestPoolOverMaxActive(t *testing.T) {
	maxActive := 10
	pool := newPool(2, maxActive)
	defer pool.Close()

	ctx := context.Background()
	for i := 0; i < maxActive; i++ {
		_, err := pool.Get(ctx)
		assert.Nil(t, err)
	}

	assert.Equal(t, pool.active, maxActive)

	_, err := pool.Get(ctx)
	assert.Equal(t, err, ErrPoolLimit)
}

func TestPoolPut(t *testing.T) {
	pool := newPool(5, 10)
	defer pool.Close()

	pc, err := pool.Get(context.Background())
	assert.Nil(t, err)

	assert.Equal(t, pool.active, 1)
	assert.Equal(t, pool.idle.count, 0)

	assert.Nil(t, pc.Close())
	assert.Equal(t, pool.active, 1)
	assert.Equal(t, pool.idle.count, 1)
}

func TestPoolMaxIdel(t *testing.T) {
	pool := &ConnectionPool{
		Dial: func(context.Context) (net.Conn, error) {
			return net.Dial(network, address)
		},
		MaxIdle:   2,
		MaxActive: 10,
	}
	defer pool.Close()

	ctx := context.Background()

	pc1, err := pool.Get(ctx)
	assert.Nil(t, err)
	pc2, err := pool.Get(ctx)
	assert.Nil(t, err)
	pc3, err := pool.Get(ctx)
	assert.Nil(t, err)

	assert.Equal(t, pool.active, 3)
	assert.Equal(t, pool.idle.count, 0)

	assert.Nil(t, pc1.Close())
	assert.Nil(t, pc2.Close())
	assert.Nil(t, pc3.Close())
	assert.Equal(t, pool.active, 2)
	assert.Equal(t, pool.idle.count, 2)

	assert.Equal(t, 2, pool.idle.count)
}

func TestPoolIdleTimeout(t *testing.T) {
	pool := newPool(5, 10)
	pool.IdleTimeout = time.Millisecond * 200
	defer pool.Close()

	ctx := context.Background()
	pcs := []*PoolConn{}
	for i := 0; i < 10; i++ {
		pc, err := pool.Get(ctx)
		assert.Nil(t, err)
		pcs = append(pcs, pc)
	}

	for _, pc := range pcs {
		assert.Nil(t, pc.Close())
	}

	assert.Equal(t, pool.idle.count, 5)
	assert.Equal(t, pool.active, 5)
	time.Sleep(time.Millisecond * 500)
	assert.Equal(t, pool.idle.count, 0)
	assert.Equal(t, pool.active, 0)
	pc, err := pool.Get(ctx)
	assert.Nil(t, err)
	assert.Equal(t, pool.idle.count, 0)
	assert.Equal(t, pool.active, 1)

	pc.Close()
	assert.Equal(t, pool.idle.count, 1)
	assert.Equal(t, pool.active, 1)
}

func TestPoolReUseConn(t *testing.T) {
	pool := newPool(5, 10)
	defer pool.Close()

	ctx := context.Background()
	for i := 0; i < 2; i++ {
		pc1, err := pool.Get(ctx)
		assert.Nil(t, err)
		pc2, err := pool.Get(ctx)
		assert.Nil(t, err)
		pc1.Close()
		pc2.Close()
	}
	assert.Equal(t, pool.idle.count, 2)
	assert.Equal(t, pool.active, 2)
}

func TestPoolMaxLifeTime(t *testing.T) {
	pool := &ConnectionPool{
		Dial: func(context.Context) (net.Conn, error) {
			return net.Dial(network, address)
		},
		MaxIdle:   5,
		MaxActive: 10,
	}

	ctx := context.Background()
	pcs := []*PoolConn{}
	for i := 0; i < 10; i++ {
		pc, err := pool.Get(ctx)
		assert.Nil(t, err)
		pcs = append(pcs, pc)
	}

	for _, pc := range pcs {
		assert.Nil(t, pc.Close())
	}

	assert.Equal(t, pool.idle.count, 5)
	assert.Equal(t, pool.active, 5)

	pool.MaxConnLifetime = time.Millisecond * 400
	pool.RegisterChecker(time.Millisecond*50, pool.defaultChecker)

	time.Sleep(time.Second * 2)

	assert.Equal(t, pool.idle.count, 0)
	assert.Equal(t, pool.active, 0)

	pc, err := pool.Get(ctx)
	assert.Nil(t, err)
	assert.Equal(t, pool.idle.count, 0)
	assert.Equal(t, pool.active, 1)

	pc.Close()
	assert.Equal(t, pool.idle.count, 1)
	assert.Equal(t, pool.active, 1)
	pool.Close()
}

func backgroundPoolGet(p *ConnectionPool, n int) chan error {
	ctx := context.Background()
	errs := make(chan error, n)
	for i := 0; i < cap(errs); i++ {
		go func() {
			c, err := p.Get(ctx)
			if c != nil {
				c.Close()
			}
			errs <- err
		}()
	}

	return errs
}

func TestPoolWait(t *testing.T) {
	pool := newPool(5, 10)
	pool.Wait = true
	defer pool.Close()

	ctx := context.Background()
	pcs := []*PoolConn{}
	for i := 0; i < 10; i++ {
		pc, err := pool.Get(ctx)
		assert.Nil(t, err)
		pcs = append(pcs, pc)
	}

	errs := backgroundPoolGet(pool, 10)
	for _, pc := range pcs {
		assert.Nil(t, pc.Close())
	}

	timeout := time.After(2 * time.Second)
	for i := 0; i < cap(errs); i++ {
		select {
		case err := <-errs:
			assert.Nil(t, err)
		case <-timeout:
			t.Fatalf("timeout waiting for blocked goroutine %d", i)
		}
	}
}

func TestPoolWaitIdleTimeout(t *testing.T) {
	pool := newPool(5, 10)
	pool.Wait = true
	pool.IdleTimeout = time.Millisecond * 200
	defer pool.Close()

	ctx := context.Background()
	pcs := []*PoolConn{}
	for i := 0; i < 10; i++ {
		pc, err := pool.Get(ctx)
		assert.Nil(t, err)
		pcs = append(pcs, pc)
	}

	for _, pc := range pcs {
		assert.Nil(t, pc.Close())
	}

	time.Sleep(time.Millisecond * 500)
	timeout := time.After(1 * time.Second)
	errs := backgroundPoolGet(pool, 10)
	for i := 0; i < cap(errs); i++ {
		select {
		case err := <-errs:
			assert.Nil(t, err)
		case <-timeout:
			t.Fatalf("timeout waiting for blocked goroutine %d", i)
		}
	}
}

func TestPoolWaitMaxLifeTime(t *testing.T) {
	pool := newPool(5, 10)
	pool.Wait = true
	pool.MaxConnLifetime = time.Millisecond * 200
	defer pool.Close()

	ctx := context.Background()
	pcs := []*PoolConn{}
	for i := 0; i < 10; i++ {
		pc, err := pool.Get(ctx)
		assert.Nil(t, err)
		pcs = append(pcs, pc)
	}

	for _, pc := range pcs {
		assert.Nil(t, pc.Close())
	}

	time.Sleep(time.Millisecond * 500)
	timeout := time.After(1 * time.Second)
	errs := backgroundPoolGet(pool, 10)
	for i := 0; i < cap(errs); i++ {
		select {
		case err := <-errs:
			assert.Nil(t, err)
		case <-timeout:
			t.Fatalf("timeout waiting for blocked goroutine %d", i)
		}
	}
}

func TestPoolDialErr(t *testing.T) {
	pool := newPool(5, 10)
	defer pool.Close()

	ctx := context.Background()
	_, err := pool.Get(ctx)
	assert.Nil(t, err)
	assert.Equal(t, pool.active, 1)
	assert.Equal(t, pool.idle.count, 0)
	pool.Dial = func(context.Context) (net.Conn, error) {
		return nil, errors.New("dial error")
	}
	_, err = pool.Get(ctx)
	assert.NotNil(t, err)
	assert.Equal(t, pool.active, 1)
	assert.Equal(t, pool.idle.count, 0)
}

func TestPoolWaitDialErr(t *testing.T) {
	pool := newPool(5, 10)
	pool.Wait = true
	defer pool.Close()

	ctx := context.Background()
	for i := 0; i < 9; i++ {
		_, err := pool.Get(ctx)
		assert.Nil(t, err)
	}

	assert.Equal(t, pool.active, 9)
	assert.Equal(t, pool.idle.count, 0)
	pool.Dial = func(context.Context) (net.Conn, error) {
		return nil, errors.New("dial error")
	}
	_, err := pool.Get(ctx)
	assert.NotNil(t, err)
	pool.Dial = func(context.Context) (net.Conn, error) {
		return net.Dial(network, address)
	}
	timeout := time.After(1 * time.Second)
	errs := backgroundPoolGet(pool, 1)
	for i := 0; i < cap(errs); i++ {
		select {
		case err := <-errs:
			assert.Nil(t, err)
		case <-timeout:
			t.Fatalf("timeout waiting for blocked goroutine %d", i)
		}
	}
	assert.Equal(t, pool.active, 10)
	assert.Equal(t, pool.idle.count, 1)
}

func TestPoolConnRead(t *testing.T) {
	pool := newPool(2, 10)
	ctx := context.Background()
	pc, err := pool.Get(ctx)
	assert.Nil(t, err)
	pc.Read(nil)
	pc.closed = true
	_, err = pc.Read(nil)
	assert.Equal(t, err, ErrConnClosed)
}

func TestPoolConnWrite(t *testing.T) {
	pool := newPool(2, 10)
	ctx := context.Background()
	pc, err := pool.Get(ctx)
	assert.Nil(t, err)
	pc.Write(nil)
	pc.closed = true
	_, err = pc.Write(nil)
	assert.Equal(t, err, ErrConnClosed)
}

func TestPoolConnectionGet(t *testing.T) {
	pool := NewConnectionPool()
	conn, err := pool.Get(network, address, time.Second)
	assert.Nil(t, err)
	assert.Nil(t, conn.Close())
}

func BenchmarkPoolGet(b *testing.B) {
	ctx := context.Background()
	p := newPool(1, 0)
	defer p.Close()
	c, err := p.Get(ctx)
	if err != nil {
		panic(err)
	}
	if err := c.Close(); err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		c, err = p.Get(ctx)
		if err != nil {
			panic(err)
		}
		if err := c.Close(); err != nil {
			panic(err)
		}
	}
}

func BenchmarkGetPoolParallel(b *testing.B) {
	b.StopTimer()
	ctx := context.Background()
	pool := newPool(10, 0)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pc, err := pool.Get(ctx)
			if err != nil {
				panic(err)
			}
			if err := pc.Close(); err != nil {
				panic(err)
			}
		}
	})
}

func TestPoolGetWithFramerBuilder(t *testing.T) {
	pool := NewConnectionPool()

	conn, err := pool.Get(network, address, time.Second, WithFramerBuilder(&emptyFramerBuilder{}))
	assert.Nil(t, err)

	fr, ok := conn.(codec.Framer)
	assert.True(t, ok)
	_, err = fr.ReadFrame()
	assert.Nil(t, err)
}
