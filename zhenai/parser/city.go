package parser

import (
	"GoSpider/engine"
	"regexp"
)

var prifileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*">([^<]+)</a>`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

func ParseCity(contents []byte) engine.ParseResult {
	matchs := prifileRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		name := m[2]
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			func(c []byte) engine.ParseResult {
				return ParseProFile(c, string(name))
			},
		})
	}

	matchs = cityUrlRe.FindAllSubmatch(contents, -1)
	for _, m := range matchs {
		result.Requests = append(result.Requests,
			engine.Request{
				Url:        string(m[1]),
				ParserFunc: ParseCity,
			})
	}
	return result
}
