package parser

import (
	"regexp"
	"GoSpider/engine"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

var idUrlRe = regexp.MustCompile(`/article/([\d]+).html`)

var meiziListPageRe=regexp.MustCompile(`(http://www.mzitu.com/.+/page/[\d]+/)`)

func ParseMeizi(contents []byte,_ string) engine.ParseResult {

	result := engine.ParseResult{}

	doc,err:=goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}

	//查找到id为pins
	doc.Find("#pins").Find("li").Each(func(i int, s *goquery.Selection) {
		meiziUrl,_:=s.Find("a").Eq(0).Attr("href")

		result.Requests = append(result.Requests, engine.Request{
			meiziUrl,
			ParseMeiziDteail,
		})
	})

	matchs:=meiziListPageRe.FindAllSubmatch(contents,-1);
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			ParseMeizi,
		})
	}

	return result
}



