package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)



func ParseColumnDetail(contents []byte,url string) engine.ParseResult {
	result := engine.ParseResult{}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}

	var items []engine.Item

	block:=doc.Find("div.article-wrap").Eq(0)

	//标题
	title:=doc.Find("title").Eq(0).Text()
	//图片
	imgSrc,_:=block.Find("div.article-img-box").Eq(0).Find("img").Eq(0).Attr("src")
	//作者
	author:=block.Find("span.author-name").Eq(0).Find("a").Eq(0).Text()
	//时间
	time:=block.Find("span.article-time.pull-left").Eq(0).Text()

	//内容容器
	contentBlock:=block.Find("div.article-content-wrap").Eq(0)

	//内容Html
	content,err:=contentBlock.Html()

	if err != nil {
		return result
	}

	new := modle.News{}
	new.Title=title
	new.Url=url
	new.ImgSrc=imgSrc
	new.Author=author
	new.TimeGap=time
	new.Content=content

	//	Each(func(i int, s *goquery.Selection) {
	//	id, _ := s.Attr("data-aid")
	//	a := s.Find("div.mod-thumb.pull-left ").Eq(0).Find("a").Eq(0)
	//	title, _ := a.Attr("title")
	//	url, _ := a.Attr("href")
	//	imgSrc, _ := a.Find("img").Eq(0).Attr("src")
	//	desc := s.Find("div.mob-ctt.channel-list-yh").Eq(0).Find("div").Last().Text()
	//
	//	new := modle.News{}
	//	new.Title = title
	//	new.Url = HOST + url
	//	new.ImgSrc = imgSrc
	//	new.Desc = desc
	//
	//	items = append(items, engine.Item{
	//		Type:    "huxiu",
	//		Id:      id,
	//		Payload: new,
	//	})
	//})

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
