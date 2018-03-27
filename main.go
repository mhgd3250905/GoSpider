package main

import (
	"GoSpider/zhenai/parser"
	"GoSpider/engine"
	"GoSpider/scheduler"
	"GoSpider/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})

	//仅仅爬取上海
	//e.Run(engine.Request{
	//	"http://www.zhenai.com/zhenghun/shanghai",
	//	parser.ParseCity,
	//})

}
