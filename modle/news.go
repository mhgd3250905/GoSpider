package modle

import "encoding/json"

//虎嗅数据modle
type News struct {
	Title   string
	Url     string
	ImgSrc  string
	TimeGap string
	Desc    string
	Content string
	Author  string
	//Column  string
}

func FromJsonObjNews(o interface{}) (News, error) {
	var news News
	s, err := json.Marshal(o)
	if err != nil {
		return news, err
	}
	err = json.Unmarshal(s, &news)
	return news, err
}
