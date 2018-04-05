package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

/*
1->url
2->imgSrc
3->title
4->desc
*/

/*
<a class="fn-clear" href="/article/285222.html" target="_blank" title="《堡垒之夜》席卷之下的美国学校">
                <img class="fn-left" src="http://img.chuapp.com//wp-content/Picture/2018-04-04/5ac4b0c4048df.jpg?imageView2/5/w/390/h/219" width="390" height="219" alt="" />
                <dl class="fn-left">
                    <dd class="fn-clear"><span class="fn-left"><em>等等</em>4小时前</span><span class="fn-right">0条评论</span></dd>
                    <dt>《堡垒之夜》席卷之下的美国学校</dt>
                    <dd>自《我的世界》以来，《堡垒之夜》成了另一款让孩子们如此痴迷的游戏。</dd>
                </dl>
            </a>
*/
var blockRe = regexp.MustCompile(`<a class="fn-clear" href="(/article/[\d]+\.html)" target="_blank" title="(.+)">
                <img class="fn-left" src="([^\"]+)" width="[\d]+" height="[\d]+" alt="" />
                <dl class="fn-left">
                    <dd class="fn-clear"><span class="fn-left"><em>[^<]+</em>[^<]+</span><span class="fn-right">[^<]+</span></dd>
                    <dt>[^<]+</dt>
                    <dd>([^<]+)</dd>
                </dl>
            </a>`)

var idUrlRe = regexp.MustCompile(`/article/([\d]+).html`)

func ParseChule(contents []byte, _ string) engine.ParseResult {

	result := engine.ParseResult{}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(contents))
	if err != nil {
		return result
	}

	divContainer:=doc.Find("div .category-list").Eq(0)

	divContainer.Find("a .fn-clear").Has("title").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Html())
	})



	matchs := blockRe.FindAllSubmatch(contents, -1)

	var items []engine.Item
	for i := 0; i < len(matchs); i++ {
		new := modle.News{}
		new.Url = string(matchs[i][1])
		new.Title = HOST + string(matchs[i][2])
		new.ImgSrc = string(matchs[i][3])
		//new.TimeGap = string(matchs[i][5])
		new.Desc = string(matchs[i][4])
		//new.Column = column

		//fmt.Printf("block:%s ,title:%s ,url:%s ,column:%s\n",matchs[i],new.Title,new.Url,column)

		items = append(items, engine.Item{
			Type:    "huxiu",
			Id:      extractString([]byte(new.Url), idUrlRe),
			Payload: new,
		})
	}

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
