package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"GoSpider/engine"
	"github.com/pkg/errors"
)

func ItemSaver(index string) (itemChan chan engine.Item,err error) {
	client, err := elastic.NewClient()
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %+v", itemCount, item)
			itemCount++

			err := save(client,index, item)
			if err != nil {
				log.Printf("Item saver:error saving item %v : %v", item, err)
			}
		}
	}()
	return out, nil
}

//save item
func save(client *elastic.Client,index string, item engine.Item) (err error) {

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	//save data -> index
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil

}
