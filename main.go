package main

import (
	"GoSpider/zhenai/parser"
	"GoSpider/engine"
	"GoSpider/scheduler"
	"GoSpider/persist"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
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
