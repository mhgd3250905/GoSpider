package persist

import (
	"testing"
	"GoSpider/modle"
	"github.com/olivere/elastic"
	"context"
	"log"
	"encoding/json"
	"GoSpider/engine"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: modle.Profile{
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
		},
	}

	const index = "dating_test"

	//todo Try to start up elastic search
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	err = save(client,index,expected)
	if err != nil {
		panic(err)
	}



	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	log.Printf("%s", resp.Source)

	var actual engine.Item
	err = json.Unmarshal(*resp.Source, &actual)
	if err != nil {
		panic(err)
	}

	actualProfile,err:=modle.FromJsonObj(actual.Payload)
	actual.Payload=actualProfile

	if actual != expected {
		t.Errorf("got %v ; expecter %v", actual, expected)
	}

}
