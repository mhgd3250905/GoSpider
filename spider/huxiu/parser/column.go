package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)



func ParseColumn(contents []byte, url string, column string) engine.ParseResult {
	result := engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}

	var items []engine.Item

	doc.Find("div.mod-info-flow").Eq(0).Find("div.mod-b.mod-art.clearfix").
		Each(func(i int, s *goquery.Selection) {
		id, _ := s.Attr("data-aid")
		a := s.Find("div.mod-thumb.pull-left ").Eq(0).Find("a").Eq(0)
		title, _ := a.Attr("title")
		url, _ := a.Attr("href")
		imgSrc, _ := a.Find("img").Eq(0).Attr("src")
		desc := s.Find("div.mob-ctt.channel-list-yh").Eq(0).Find("div").Last().Text()

		new := modle.News{}
		new.Title = title
		new.Url = HOST + url
		new.ImgSrc = imgSrc
		new.Desc = desc

		items = append(items, engine.Item{
			Type:    "huxiu",
			Id:      id,
			Payload: new,
		})
	})

	result = engine.ParseResult{
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
