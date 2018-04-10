package parser

import (
	"GoSpider/engine"
	"regexp"
)

const HOST = `http://www.mzitu.com/`
var meiziListRe = regexp.MustCompile(`href="(http://www.mzitu.com/[^\d">]+)"`)

func ParseMeiziList(contents []byte, host string) engine.ParseResult {
	matchs := meiziListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			ParseMeizi,
		})
	}
	return result
}

