package parser

import (
	"GoSpider/engine"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*">([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		name := m[2]
		result.Items = append(result.Items, "User "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			func(c []byte) engine.ParseResult {
				return ParseProFile(c, string(name))
			},
		})
	}
	return result
}
