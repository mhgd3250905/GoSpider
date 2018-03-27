package persist

import (
	"testing"
	"GoSpider/modle"
	"github.com/olivere/elastic"
	"context"
	"log"
	"encoding/json"
)

func TestSave(t *testing.T) {
	expected := modle.Profile{
		Age:        34,
		Height:     162,
		Weight:     57,
		Income:     "3001-5000元",
		Gender:     "女",
		Name:       "安静的雪",
		Xinzuo:     "牡羊座",
		Occupation: "人事/行政",
		Marriage:   "离异",
		House:      "已购房",
		Hukou:      "山东菏泽",
		Education:  "大学本科",
		Car:        "未购车",
	}

	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	resp,err:=client.Get().
		Index("dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	log.Printf("%s",resp.Source)

	var actual modle.Profile
	err=json.Unmarshal(*resp.Source,&actual)
	if err != nil {
		panic(err)
	}

	if actual!=expected {
		t.Errorf("got %v ; expecter %v",actual,expected)
	}

}
