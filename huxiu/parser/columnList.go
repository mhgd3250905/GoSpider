package parser

import (
	"GoSpider/engine"
	"regexp"
)

const host  = `https://www.huxiu.com/`
const columnListRe = `<a href="(/channel/[\d]+.html)" [^>]+>([^<]+)</a>`

func ParseColumnList(contents []byte,host string) engine.ParseResult {
	re := regexp.MustCompile(columnListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			host+string(m[1]),
			ParseColumn,
		})
	}
	return result
}

