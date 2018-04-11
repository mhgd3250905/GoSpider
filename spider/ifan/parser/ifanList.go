package parser

import (
	"regexp"
	"GoSpider/engine"
	"fmt"
)

const HOST = `https://www.ifanr.com/`

/**
1.categoryUrl
2.page
 */
var categoryUrlModle=`%s?page=%d&pajax=1&post_id__lt=9999999`
/**
1->url
 */
var ifanCategoryRe = regexp.MustCompile(`<li class="menu-wrap__item menu-wrap-seprator"><a href="(http://www.ifanr.com/category/[^"]+)">[^<]+</a></li>`)

func ParseIfanList(contents []byte, host string) engine.ParseResult {
	//获取category

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