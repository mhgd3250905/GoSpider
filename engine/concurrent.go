package engine

import (
	"log"
	"GoSpider/modle"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
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

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("id: %s,title: %s",item,item.Payload.(modle.HuxiuNews).Title)
			//save item
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			//判断是否重复
			//if isDupllication(request.Url){
			//	continue
			//}
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

//去重复
var visitedUrls=make(map[string]bool)

func isDupllication(url string) bool{
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url]=true
	return false
}
