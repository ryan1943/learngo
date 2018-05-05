package engine

import (
	"learngo/concurrencycrawler/fetcher"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request //生成worker chan
	Run()
}

type ReadyNotifier interface {
	//发送worker chan 到scheduler的workerChan，它的类型是 chan chan engine.Request
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	//生成相应数目的worker chan，并发送到scheduler的worker队列去
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r) //发送request到scheduler的request队列去
	}
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item: #%d: %v", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			//发送到worker队列去
			ready.WorkerReady(in)
			//接收到这里传来的数据 activeWorker <- activeRequest, 可能要等待
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParserFunc(body), nil
}
