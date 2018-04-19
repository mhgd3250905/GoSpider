package parser

import (
	"GoSpider/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

func ParseColumn(contents []byte, url string, column string) engine.ParseResult {
	result := engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}

	doc.Find("div.mod-info-flow").Eq(0).Find("div.mod-b.mod-art.clearfix").
		Each(func(i int, s *goquery.Selection) {
		a := s.Find("div.mod-thumb.pull-left ").Eq(0).Find("a").Eq(0)
		url, _ := a.Attr("href")

		result.Requests = append(result.Requests, engine.Request{
			HOST + url,
			ColumnDetialParser(HOST + url),
		})
	})

	return result
}

func ColumnDetialParser(column string) engine.ParserFunc {
	return func(c []byte, _ string) engine.ParseResult {
		return ParseColumnDetail(c, column)
	}
}
