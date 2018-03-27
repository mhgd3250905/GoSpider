package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return out
}

//save item
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient()
	if err != nil {
		return "", err
	}

	//save data -> index
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "", err
	}

	return resp.Id, nil

}
