package parser

import (
	"GoSpider/engine"
	"GoSpider/modle"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

type TopicBean struct {
	Code        int `json:"code"`
	Timestamp   int `json:"timestamp"`
	TimestampRt int `json:"timestamp_rt"`
	Data struct {
		TotalCount int `json:"total_count"`
		Page       int `json:"page"`
		PageSize   int `json:"page_size"`
		Items []struct {
			ID int `json:"id"`
			//Summary          string `json:"summary"`
			//ProjectID        string `json:"project_id"`
			//TagID            int    `json:"tag_id"`
			//ViewsCount       string `json:"views_count"`
			//MobileViewsCount string `json:"mobile_views_count"`
			//AppViewsCount    string `json:"app_views_count"`
			//MonographicID    int    `json:"monographic_id"`
			//DomainID         string `json:"domain_id"`
			//GoodsID          string `json:"goods_id"`
			//IsTovc           string `json:"is_tovc"`
			//IsFree           string `json:"is_free"`
			//ColumnName       string `json:"column_name"`
			Title string `json:"title"`
			Cover string `json:"cover"`
			//TemplateInfo struct {
			//	TemplateType        string   `json:"template_type"`
			//	TemplateTitle       string   `json:"template_title"`
			//	TemplateTitleIsSame bool     `json:"template_title_isSame"`
			//	TemplateCover       []string `json:"template_cover"`
			//} `json:"template_info"`
			//PublishedAt    time.Time   `json:"published_at"`
			//ColumnID       string      `json:"column_id"`
			//UserID         string      `json:"user_id"`
			//ExtractionTags string      `json:"extraction_tags"`
			//UserInfo       string      `json:"user_info"`
			//Highlight      interface{} `json:"highlight"`
			//Type           string      `json:"_type"`
			//Score          interface{} `json:"_score"`
			//FavouriteNum   int         `json:"favourite_num"`
		} `json:"items"`
	} `json:"data"`
}

/*
1->id
*/
var HOST = `http://36kr.com/p/%s.html`

func ParseTopic(contents []byte, _ string) engine.ParseResult {
	var topic TopicBean
	json.Unmarshal(contents, &topic)

	matchs := topic.Data.Items

	var items []engine.Item
	var result engine.ParseResult
	for i := 0; i < len(matchs); i++ {
		new := modle.News{}
		new.Title = matchs[i].Title
		new.Url = fmt.Sprintf(HOST, strconv.Itoa(matchs[i].ID))
		new.ImgSrc = matchs[i].Cover
		//new.Column = column

		//fmt.Printf("block:%s ,title:%s ,url:%s ,column:%s\n",matchs[i],new.Title,new.Url,column)

		items = append(items, engine.Item{
			Type:    "ke36",
			Id:      strconv.Itoa(matchs[i].ID),
			Payload: new,
		})
	}

	result.Items = items
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
