package parser

import (
	"GoSpider/engine"
	"fmt"
	"time"
)

/*
1->topId
2->时间戳
*/
var topicUrl=`http://36kr.com/api/search-column/%d?per_page=20&page=1&_=%d`

func ParseTopicList(contents []byte, _ string) engine.ParseResult {
	var topicIds = []int{23, 221, 225, 218}
	result := engine.ParseResult{}
	for _, topicId := range topicIds {
		result.Requests = append(result.Requests, engine.Request{
			fmt.Sprintf(topicUrl,topicId,time.Now().Unix()),
			ParseTopic,
		})
	}
	return result
}
