package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
	"github.com/PuerkitoBio/goquery"
	"bytes"
	"time"
	"strconv"
)

/**
1->imgSrc
 */
var imgSrcRe = regexp.MustCompile(`background-image: url\((http://cdn.pingwest.com/wp-content/uploads/[^-]+)-620x0\);`)

func ParsePingwest(contents []byte, host string) engine.ParseResult {
	//获取category

	result := engine.ParseResult{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}
	var items []engine.Item
	doc.Find("div.news-item").Each(func(i int, s *goquery.Selection) {
		a := s.Find("h2").Eq(0).Find("a").Eq(0)
		url, _ := a.Attr("href")
		title := a.Text()
		desc := s.Find("div.des").Eq(0).Text()
		imgSrcStr, _ := s.Find("div.news-thumb").Eq(0).Attr("style")
		imgSrc := imgSrcRe.FindAllStringSubmatch(imgSrcStr, -1)[0][1]

		new := modle.News{}

		new.Url = url
		new.Title = title
		new.ImgSrc = imgSrc
		new.Desc = desc

		items = append(items, engine.Item{
			Type:    "pingwest",
			Id:      strconv.FormatInt(time.Now().Unix(), 10),
			Payload: new,
		})
	})

	result = engine.ParseResult{
		Items: items,
	}
	return result
}
