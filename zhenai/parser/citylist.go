package parser

import (
	"GoSpider/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			ParseCity,
		})
	}
	return result
}
