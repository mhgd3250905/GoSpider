package main

import (

	"GoSpider/zhenai/parser"
	"GoSpider/engine"
	"GoSpider/scheduler"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount:10,
	}

	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})

}
