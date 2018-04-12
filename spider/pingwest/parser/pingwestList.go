package parser

import (
	"GoSpider/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

const HOST = `https://www.ifanr.com/`

/**
1.categoryUrl
2.page
 */
/**
1->url
 */

func ParsePingwestList(contents []byte, host string) engine.ParseResult {
	//获取category
	result := engine.ParseResult{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}
	doc.Find("ul.dropdown-menu").Eq(0).Find("li").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find("a").Eq(0).Attr("href")
		result.Requests = append(result.Requests, engine.Request{
			url,
			ParsePingwest,
		})
	})

	return result
}
