package parser

import (
	"regexp"
	"GoSpider/engine"
	"GoSpider/modle"
)


/**
1->id
2->imgSrc
3->url
4->title
 */
var ifanThumbnailRe = regexp.MustCompile(`<div class="c-card c-card-article c-card-article--thumbnail" data-post-id="([\d]+)">
            <label class="c-card-article__label"><a class="c-card-article__label__link" href="http://www.ifanr.com/category/[^"]+" target="_blank">[^<]+</a></label>
            <div class="c-card-article__thumbnail" style="background-image:url\('(https://images.ifanr.cn/wp-content/uploads/[^\.]+.jpg)![\d]+'\)">
              <h1 class="c-card-article__title">
                <a class="c-card-article__link" href="(http://www.ifanr.com/[\d]+)" target="_blank">([^<]+)</a>
              </h1>
            </div>
            <a class="c-card-article__thumbnail__link" href="http://www.ifanr.com/[\d]+" target="_blank"></a>
          </div>`)

/**
1->id
2->imgSrc
3->url
4->title
 */
var ifanRe=regexp.MustCompile(`div class="article-item article-item--card" data-post-id="([\d]+)">
            <div class="article-image cover-image" style="background-image: url\('(https://images.ifanr.cn/wp-content/uploads/[^\.]+.jpg!260)'\);"></div>
            <a href="http://www.ifanr.com/category/[^"]+" class="article-label" target="_blank">[^<]+</a>
            <a href="(http://www.ifanr.com/[\d]+)" class="article-link cover-block" target="_blank"></a>
            <h3>([^<]+)</h3>
            <time>[^<]+</time>
            <div class="article-meta" data-post-id="[\d]+">
              <span class="ifanrx-like like-count js-article-like-count">-</span>
              <a class="text-link" href="http://www.ifanr.com/[\d]+#article-comments" target="_blank"><span class="ifanrx-reply comment-count">[^<]+</span></a>
            </div>
          </div>`)



func ParseIfan(contents []byte, host string) engine.ParseResult {
	//获取category

	result := engine.ParseResult{}

	matchs := ifanThumbnailRe.FindAllSubmatch(contents, -1)

	var items []engine.Item
	for _, m := range matchs {
		new := modle.News{}


		new.Url = string(m[3])
		new.Title = string(m[4])
		new.ImgSrc = string(m[2])
		//new.Desc = string(m[3])

		items = append(items, engine.Item{
			Type:    "ifan",
			Id:      string(m[1]),
			Payload: new,
		})
	}

	matchs = ifanRe.FindAllSubmatch(contents, -1)

	for _, m := range matchs {
		new := modle.News{}


		new.Url = string(m[3])
		new.Title = string(m[4])
		new.ImgSrc = string(m[2])
		//new.Desc = string(m[3])

		items = append(items, engine.Item{
			Type:    "ifan",
			Id:      string(m[1]),
			Payload: new,
		})
	}

	result=engine.ParseResult{
		Items:items,
	}
	return result
}