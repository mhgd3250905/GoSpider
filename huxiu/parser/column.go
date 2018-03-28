package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
)

var titleRe = regexp.MustCompile(`<a class="transition" title="(.+)"
                   href="(/article/[\d]+\.html)" [^>]+>`)
var urlRe = regexp.MustCompile(`<a class="transition" title=".+"
                   href="(/article/[\d]+\.html)" [^>]+>`)
var idUrlRe = regexp.MustCompile(`<a class="transition" title=".+"
                   href="/article/([\d]+)\.html" [^>]+>`)

func ParseColumn(contents []byte,_ string) engine.ParseResult {

	new:=modle.HuxiuNews{}

	new.Title = extractString(contents, titleRe)
	new.Url = extractString(contents, urlRe)

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Type:    "huxiu",
				Id:      extractString([]byte(new.Url), idUrlRe),
				Payload: new,
			},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}