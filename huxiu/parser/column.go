package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
)

/*
1->title
2->url
3->imgSrc
4->time
5->desc
*/
var blockRe = regexp.MustCompile(`<div class="mod-thumb pull-left ">
                                <a class="transition" title="([^"]+)"
                   href="(/article/[\d]+.html)" target="_blank">
                                            <img class="lazy"
                             data-original="(https://img.huxiucdn.com/article/cover/[^"]+)"
                             [^>]+>
                                    </a>
            </div>
            <div class="mob-ctt channel-list-yh">
                <h2>
                    <a href="/article/[\d]+.html" class="transition msubstr-row2" target="_blank">[^<]+</a>
                </h2>
                <div class="mob-author">
                                            <div class="author-face">
                            <a [^>]+><img [^>]+></a>
                        </div>
                        <a [^>]+>
                            <span class="author-name">[^<]+</span>
                        </a>
                        <a href="/vip" target="_blank">[^<]*</a>
                        (<i class="i-icon icon-auth2" title="虎嗅认证作者"></i>){0,1}                                        <span class="time">([^<]+)</span>
                    <i class="icon icon-cmt"></i><em>[^<]+</em>
                    <i class="icon icon-fvr"></i><em>[^<]+</em>
                </div>

                <div class="mob-sub">([^<]+)</div>
            </div>`)

var idUrlRe = regexp.MustCompile(`/article/([\d]+).html`)

func ParseColumn(contents []byte, url string, column string) engine.ParseResult {

	matchs := blockRe.FindAllSubmatch(contents, -1)

	var items []engine.Item
	for i := 0; i < len(matchs); i++ {
		new := modle.HuxiuNews{}
		new.Title = string(matchs[i][1])
		new.Url = HOST + string(matchs[i][2])
		new.ImgSrc = string(matchs[i][3])
		new.TimeGap = string(matchs[i][5])
		new.Desc = string(matchs[i][6])
		//new.Column = column

		//fmt.Printf("block:%s ,title:%s ,url:%s ,column:%s\n",matchs[i],new.Title,new.Url,column)

		items = append(items, engine.Item{
			Type:    "huxiu",
			Id:      extractString([]byte(new.Url), idUrlRe),
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
