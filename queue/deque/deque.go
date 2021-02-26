package deque

import (
	"code_tpl_go/util"
	"context"
	"github.com/siddontang/go-log/log"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type ProducerFunc func(int) ([]string, error)
type ConsumerFunc func(string) error

type Deque struct {
	chann           chan string
	producer        ProducerFunc
	consumer        ConsumerFunc
	consumeLimiter  *rate.Limiter
	produceBatchNum int
}

func NewDeque(producerFunc ProducerFunc, consumerFunc ConsumerFunc, produceBatchNum int, ConsumeQpm int) *Deque {
	limit := rate.Every(time.Minute / time.Duration(ConsumeQpm))
	limiter := rate.NewLimiter(limit, 1)
	return &Deque{
		chann:           make(chan string, produceBatchNum*2),
		producer:        producerFunc,
		consumer:        consumerFunc,
		produceBatchNum: produceBatchNum,
		consumeLimiter:  limiter,
	}
}

func (deque *Deque) Run() {
	go deque.produceLoop()
	go deque.consumeLoop()
}

func (deque *Deque) produceLoop() {
	for {
		elements, err := deque.producer(deque.produceBatchNum)
		util.ReportMonitor("deque.producer-总量")
		if err != nil {
			util.ReportMonitor("deque.producer-err")
			log.Errorf("producer err:%v", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		if len(elements) == 0 {
			util.ReportMonitor("deque.producer-生产elements为空")
			log.Infof("produer elements empty")
			time.Sleep(time.Millisecond * 50)
			continue
		}
		util.ReportMonitor("deque.producer成功")
		log.Debugf("produer succ elements size:%v", len(elements))
		for _, e := range elements {
			deque.chann <- e
			util.ReportMonitor("deque.produce-one")
		}
	}
}

func (deque *Deque) consumeOne(e string, wg *sync.WaitGroup) {
	if wg != nil {
		defer wg.Done()
	}
	util.ReportMonitor("deque.consumer-总量")
	err := deque.consumer(e)
	if err != nil {
		util.ReportMonitor("deque.consumer-err")
		log.Errorf("consumer err:%v", err)
	} else {
		util.ReportMonitor("deque.consumer成功")
	}
}

func (deque *Deque) consumeLoop() {
	for {
		for e := range deque.chann {
			deque.consumeLimiter.Wait(context.Background())
			// fixme wg.Add(), wg.Done() 应该放在同一层次
			go deque.consumeOne(e, nil)
		}
	}
}

func (deque *Deque) RunSimple() {
	//一轮完全消费完毕后才开始下一轮。
	go deque.produceAndConsumeLoop()
}

func (deque *Deque) produceAndConsumeLoop() {
	for {
		elements, err := deque.producer(deque.produceBatchNum)
		util.ReportMonitor("deque.producer-总量")
		if err != nil {
			util.ReportMonitor("deque.producer-err")
			log.Errorf("producer err:%v", err)
			time.Sleep(time.Millisecond * 100)
			continue
		}
		util.ReportMonitor("deque.producer成功")

		wg := &sync.WaitGroup{}
		for _, e := range elements {
			deque.consumeLimiter.Wait(context.Background())
			util.ReportMonitor("deque.produce-one")

			wg.Add(1)
			go deque.consumeOne(e, wg)
		}
		wg.Wait() //一轮完全消费完毕后才开始下一轮
	}
}
