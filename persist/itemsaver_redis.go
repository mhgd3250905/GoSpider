package persist

import (
	"log"
	"github.com/garyburd/redigo/redis"
	"GoSpider/engine"
	"github.com/pkg/errors"
	"encoding/json"
	"GoSpider/modle"
	"time"
)

func ItemSaverRedis(index string) (itemChan chan engine.Item, err error) {
	//创建一个Redis实例
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		return nil, err
	}
	//输入账号密码
	//_, err = conn.Do("AUTH", "sk3250905")
	//if err != nil {
	//	return nil, errors.Errorf("Redis AUTH fail %v",err)
	//}
	//测试连接
	if result, err := conn.Do("ping"); result != "PONG" {
		return nil, errors.Errorf("Redis ping fail %v", err)

	}

	out := make(chan engine.Item)

	//saveMaps:=make(map[string]bool)

	go func() {
		itemCount := 0
		for {
			item := <-out

			//if saveMaps[item.Id] {
			//	fmt.Println("item 已经保存过了，跳过")
			//	continue
			//}
			//
			//saveMaps[item.Id]=true

			log.Printf("Item Saver: Got item #%d: %+v", itemCount, item)
			itemCount++

			if item.Type == "meizi" {
				news, err := modle.FromJsonObjNews(item.Payload)
				if err != nil {
					continue
				}
				time.Sleep(time.Millisecond*50)
				go DownloadFile(news.Url, news.Title, item.Id)
			}else {
				err := saveRedis(conn, index, item)
				if err != nil {
					log.Printf("Item saver:error saving item %v : %v", item, err)
				}
			}
		}
	}()
	return out, nil
}

//save item
func saveRedis(conn redis.Conn, index string, item engine.Item) (err error) {

	if err != nil {
		return err
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	id := item.Id

	itemBuf, err := json.Marshal(item)
	if err != nil {
		return err
	}

	_, err = conn.Do("ZADD", index, id, string(itemBuf))

	if err != nil {
		return err
	}

	return nil

}
