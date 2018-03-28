package main

import (
	"GoSpider/engine"
	"GoSpider/scheduler"
	"GoSpider/persist"
	"GoSpider/huxiu/parser"
)

func main() {
	itemChan, err := persist.ItemSaver("dating_profile_2")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 20,
		ItemChan:    itemChan,
	}

	e.Run(engine.Request{
		"https://www.huxiu.com/",
		parser.ParseColumnList,
	})

	//仅仅爬取上海
	//e.Run(engine.Request{
	//	"http://www.zhenai.com/zhenghun/shanghai",
	//	parser.ParseCity,
	//})

}
