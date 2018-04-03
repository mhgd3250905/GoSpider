package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
)

/*
1->url
2->imgUrl
3->title
*/
var itemRe = regexp.MustCompile(`<div class="wrap"><div class="box"><a href="(/story/[\d]+)" class="link-button"><img src="([^"]+)" class="preview-image"><span class="title">([^<]+)</span></a></div></div>`)

/*
<div class="wrap">
                        <div class="box"><a href="/story/9676894" class="link-button"><img
                                src="https://pic3.zhimg.com/v2-8b2c4ce03c630231618f7845354ebfca.jpg"
                                class="preview-image"><span class="title">武汉理工大学研究生自杀，「天然畸形」的高校师生关系脱不了干系</span></a>
                        </div>
                    </div>

*/

var idUrlRe = regexp.MustCompile(`/story/([^d]+)`)

var HOST = `https://daily.zhihu.com`

func ParseZhihuHome(contents []byte, _ string) engine.ParseResult {

	matchs := itemRe.FindAllSubmatch(contents, -1)

	var items []engine.Item
	for i := 0; i < len(matchs); i++ {
		new := modle.News{}
		new.Title = string(matchs[i][3])
		new.Url = HOST + string(matchs[i][1])
		new.ImgSrc = string(matchs[i][2])
		//new.Column = column

		//fmt.Printf("block:%s ,title:%s ,url:%s ,column:%s\n",matchs[i],new.Title,new.Url,column)

		items = append(items, engine.Item{
			Type:    "zhihu",
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
