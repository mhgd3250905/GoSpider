package parser

import (
	"GoSpider/engine"
	"regexp"
	"GoSpider/modle"
	"strconv"
)

/**
1.id-1
2.id-2
3.img-3
4.title
 */
var meiziImgRe = regexp.MustCompile(`<a href="http://www.mzitu.com/([\d]+)/([\d]+)" ><img src="(http://i.meizitu.net/.+\.jpg)" alt="([^"]+)" /></a>`)

var meiziPageRe = regexp.MustCompile(`href='(http://www.mzitu.com/[\d]+/[\d]+)'`)


func ParseMeiziDteail(contents []byte, _ string) engine.ParseResult {

	matchs := meiziImgRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	var items []engine.Item
	for i := 0; i < len(matchs); i++ {
		new := modle.News{}
		new.Title = string(matchs[i][4])
		new.Url = string(matchs[i][3])

		id_1, err := strconv.Atoi(string(matchs[i][1]))
		if err != nil {
			continue
		}
		id_2, err := strconv.Atoi(string(matchs[i][2]))
		if err != nil {
			continue
		}

		id := id_1*10000 + id_2

		items = append(items, engine.Item{
			Url:     string(matchs[i][3]),
			Type:    "meizi",
			Id:      strconv.Itoa(id),
			Payload: new,
		})
	}

	result = engine.ParseResult{
		Items: items,
	}

	pageMatchs := meiziPageRe.FindAllSubmatch(contents, -1)

	for _, m := range pageMatchs {
		result.Requests = append(result.Requests, engine.Request{
			string(m[1]),
			ParseMeiziDteail,
		})
	}

	return result
}
