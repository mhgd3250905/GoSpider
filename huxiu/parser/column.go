package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
)

var blockRe = regexp.MustCompile(`<a class="transition" title=".+"
                   href="/article/[\d]+.html" target="_blank">`)
var titleRe = regexp.MustCompile(`<a class="transition" title="(.+)"
                   href="/article/[\d]+.html" target="_blank">`)
var urlRe = regexp.MustCompile(`<a class="transition" title=".+"
                   href="(/article/[\d]+\.html)" [^>]+>`)
var idUrlRe = regexp.MustCompile(`<a class="transition" title=".+"
                   href="/article/([\d]+)\.html" [^>]+>`)

func ParseColumn(contents []byte, url string, column string) engine.ParseResult {

	matchs := blockRe.FindAll(contents, -1)

	var items []engine.Item
	for i := 0; i < len(matchs); i++ {
		new := modle.HuxiuNews{}
		new.Title = extractString([]byte(matchs[i]), titleRe)
		new.Url = HOST + extractString([]byte(matchs[i]), urlRe)
		new.Column = column

		//fmt.Printf("block:%s ,title:%s ,url:%s ,column:%s\n",matchs[i],new.Title,new.Url,column)

		items = append(items, engine.Item{
			Type:    "huxiu",
			Id:      extractString([]byte(matchs[i]), idUrlRe),
			Payload: new,
		})
	}

	result := engine.ParseResult{
		Items: items,
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
