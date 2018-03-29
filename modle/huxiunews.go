package modle

import "encoding/json"

//珍爱网数据modle
type HuxiuNews struct {
	Title  string
	Url    string
	Column string
}

func FromJsonObjHuxiu(o interface{}) (HuxiuNews, error) {
	var news HuxiuNews
	s, err := json.Marshal(o)
	if err != nil {
		return news, err
	}
	err = json.Unmarshal(s, &news)
	return news, err
}
