package main

import (
	"GoSpider/engine"
	"GoSpider/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		"http://www.zhenai.com/zhenghun",
		parser.ParseCityList,
	})

}
