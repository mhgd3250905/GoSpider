package main

import (
	"GoSpider/engine"
	"GoSpider/scheduler"
	"GoSpider/persist"
	//chule "GoSpider/spider/chule/parser"
	//huxiu "GoSpider/spider/huxiu/parser"
	//ke36 "GoSpider/spider/36ke/parser"
	//ifan "GoSpider/spider/ifan/parser"
	//pingwest "GoSpider/spider/pingwest/parser"
	huxiuDetail "GoSpider/spider/huxiuDetail/parser"
)

func main() {
	//for {
	//
	//
	//	go func() {
	//		itemChan, err := persist.ItemSaverRedis("huxiu")
	//		if err != nil {
	//			panic(err)
	//		}
	//		e := engine.ConcurrentEngine{
	//			Scheduler:   &scheduler.QueueScheduler{},
	//			WorkerCount: 20,
	//			ItemChan:    itemChan,
	//			Header: map[string]string{
	//				"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//				"Cookie":     `screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; aliyungf_tc=AQAAAOuQyXg3NA0A40KttGHnrOtzsJK1; screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; SERVERID=03a07aad3597ca2bb83bc3f5ca3decf7|1522469603|1522469119`,
	//			},
	//		}
	//
	//		e.Run(engine.Request{
	//			"https://www.huxiu.com/",
	//			huxiu.ParseColumnList,
	//		})
	//	}()
	//
	//	go func() {
	//		itemChan, err := persist.ItemSaverRedis("ke36")
	//		if err != nil {
	//			panic(err)
	//		}
	//		e := engine.ConcurrentEngine{
	//			Scheduler:   &scheduler.QueueScheduler{},
	//			WorkerCount: 20,
	//			ItemChan:    itemChan,
	//			Header: map[string]string{
	//				"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//			},
	//		}
	//
	//		e.Run(engine.Request{
	//			"https://daily.zhihu.com/", //随便给一个就行
	//			ke36.ParseTopicList,
	//		})
	//	}()
	//
	//	go func() {
	//		itemChan, err := persist.ItemSaverRedis("chule")
	//		if err != nil {
	//			panic(err)
	//		}
	//		e := engine.ConcurrentEngine{
	//			Scheduler:   &scheduler.QueueScheduler{},
	//			WorkerCount: 20,
	//			ItemChan:    itemChan,
	//			Header: map[string]string{
	//				"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//			},
	//		}
	//
	//		e.Run(engine.Request{
	//			"http://www.chuapp.com/", //随便给一个就行
	//			chule.ParseChuleList,
	//		})
	//
	//	}()
	//
	//	go func() {
	//		itemChan, err := persist.ItemSaverRedis("ifan")
	//		if err != nil {
	//			panic(err)
	//		}
	//		e := engine.ConcurrentEngine{
	//			Scheduler:   &scheduler.QueueScheduler{},
	//			WorkerCount: 20,
	//			ItemChan:    itemChan,
	//			Header: map[string]string{
	//				"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//			},
	//		}
	//
	//		e.Run(engine.Request{
	//			"https://www.ifanr.com/",
	//			ifan.ParseIfanList,
	//		})
	//	}()
	//
	//	go func() {
	//		itemChan, err := persist.ItemSaverRedis("pingwest")
	//		if err != nil {
	//			panic(err)
	//		}
	//		e := engine.ConcurrentEngine{
	//			Scheduler:   &scheduler.QueueScheduler{},
	//			WorkerCount: 20,
	//			ItemChan:    itemChan,
	//			Header: map[string]string{
	//				"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//			},
	//		}
	//
	//		e.Run(engine.Request{
	//			"http://www.pingwest.com/",
	//			pingwest.ParsePingwestList,
	//		})
	//	}()
	//
	//	fmt.Println("等待30分钟后再次获取资源...")
	//
	//	<-time.After(time.Minute * 20)
	//}

	itemChan, err := persist.ItemSaverRedis("huxiuDetail")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 20,
		ItemChan:    itemChan,
		Header: map[string]string{
			"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
			"Cookie":     `screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; aliyungf_tc=AQAAAOuQyXg3NA0A40KttGHnrOtzsJK1; screen=%7B%22w%22%3A1536%2C%22h%22%3A864%2C%22d%22%3A1.25%7D; SERVERID=03a07aad3597ca2bb83bc3f5ca3decf7|1522469603|1522469119`,
		},
	}

	e.Run(engine.Request{
		"https://www.huxiu.com/",
		huxiuDetail.ParseColumnList,
	})

	//itemChan, err := persist.ItemSaverRedis("meizi")
	//if err != nil {
	//	panic(err)
	//}
	//e := engine.ConcurrentEngine{
	//	Scheduler:   &scheduler.QueueScheduler{},
	//	WorkerCount: 20,
	//	ItemChan:    itemChan,
	//	Header: map[string]string{
	//		"User-Agent": `Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36`,
	//	},
	//}
	//
	//e.Run(engine.Request{
	//	"http://www.mzitu.com",
	//	meizi.ParseMeiziList,
	//})

	//仅仅爬取上海
	//e.Run(engine.Request{
	//	"http://www.zhenai.com/zhenghun/shanghai",
	//	parser.ParseCity,
	//})

}
