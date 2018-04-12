package parser

import (
	"regexp"
	"GoSpider/engine"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
)

const HOST = `https://www.ifanr.com/`

/**
1->url
 */
var ifanCategoryRe = regexp.MustCompile(`<li class="menu-wrap__item menu-wrap-seprator"><a href="(http://www.ifanr.com/category/[^"]+)">[^<]+</a></li>`)

func ParseIfanList(contents []byte, host string) engine.ParseResult {
	//获取category
	goquery.NewDocumentFromReader()


	matchs := ifanCategoryRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Requests = append(result.Requests, engine.Request{
			fmt.Sprintf(categoryUrlModle,string(m[1]),1),
			ParseIfan,
		})
	}
	return result
}