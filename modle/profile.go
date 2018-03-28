package modle

import (
	"encoding/json"
)

//珍爱网数据modle
type Profile struct {
	Name       string
	Gender     string
	Age        int
	Height     int
	Weight     int
	Income     string
	Marriage   string //婚况
	Education  string //教育
	Occupation string
	Hukou      string //籍贯
	Xinzuo     string
	House      string
	Car        string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}
	err = json.Unmarshal(s,&profile)
	return profile, err
}
