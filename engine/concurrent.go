package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {

	// 非缓冲型管道
	in := make(chan Request)
	out := make(chan ParseResult)
	// 将创建好的input channel传给Scheduler
	c.Scheduler.ConfigureMasterWorkerChan(in)
	// 创建WorkerCount个goroutine执行任务
	for i := 0; i < c.WorkerCount; i++ {
		createWorker(in, out)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			// 第一次运行到这里，in中并没有任务，休眠阻塞
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
