package main

import (
	"GoSpider/engine"
	"GoSpider/scheduler"
	"GoSpider/persist"
	"GoSpider/huxiu/parser"
)

func main() {
	itemChan, err := persist.ItemSaverRedis("dating_profile_2")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 20,
		ItemChan:    itemChan,
		Header: map[string]string{
			"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
			"Cookie":`screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; aliyungf_tc=AQAAAOuQyXg3NA0A40KttGHnrOtzsJK1; screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; SERVERID=03a07aad3597ca2bb83bc3f5ca3decf7|1522469603|1522469119`,
			},
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
