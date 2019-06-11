package engine

import (
	"learngo/crawler/toolkit/cmap"
)

type Processor func(Request) (ParseResult, error)

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			// 从worker的chan request中读取到数据(request),这些数据是从队列中读取到的,本身就是Request
			// worker里面的workerChan本身就是存储的Request
			// 而且都是QueuedScheduler的run函数,从request对列取出来,转移到worker对列
			// 在run函数中，req数据和worker被排队，选中的数据被分配到worker中了，所以要获取req，是从worker中获取
			request := <-in
			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	var pairRedistributor cmap.PairRedistributor
	m, err := cmap.NewConcurrentMap(10, pairRedistributor)
	if err != nil {
		panic(err)
	}

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if m.Get(r.Url) != nil {
			continue
		}

		m.Put(r.Url, struct{}{})
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		for _, request := range result.Requests {
			if m.Get(request.Url) != nil {
				continue
			}

			m.Put(request.Url, struct{}{})
			e.Scheduler.Submit(request)
		}
	}
}
