package parser

import (
	"GoSpider/engine"
	"regexp"
)

const HOST = `https://www.huxiu.com`
const columnListRe = `<a href="(/channel/[\d]+.html)" [^>]+>([^<]+)</a>`

func ParseColumnList(contents []byte, host string) engine.ParseResult {
	re := regexp.MustCompile(columnListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			HOST + string(m[1]),
			ColumnParser(string(m[2])),
		})
	}
	return result
}

func ColumnParser(column string) engine.ParserFunc {
	return func(c []byte, url string) engine.ParseResult {
		return ParseColumn(c, url, column)
	}
}
