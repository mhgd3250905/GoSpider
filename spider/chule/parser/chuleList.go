package parser

import (
	"GoSpider/engine"
	"regexp"
)

const HOST = `http://www.chuapp.com`

//<li><a href="/category/daily" target="_blank">每日聚焦</a></li>

var categoryRe = regexp.MustCompile(`<li><a href="(/category/.+)" target="_blank">[^<]+</a></li>`)
var tagIdRe = regexp.MustCompile(`<li><a href="(/tag/index/id/[\d]+.html)" target="_blank">[^<]+</a></li>`)

func ParseChuleList(contents []byte, host string) engine.ParseResult {
	//获取category
	//fmt.Println(string(contents))

	matchs := categoryRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			HOST + string(m[1]),
			ParseChule,
		})
	}

	//获取tag
	matchs = tagIdRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			HOST + string(m[1]),
			ParseChule,
		})
	}
	return result
}
