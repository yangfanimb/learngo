package persist

import (
	"context"
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"

	"github.com/olivere/elastic"
)

var logger = log.DLogger()

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.66.102:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			logger.Info("Got item #%d: %v\n", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				logger.Error("save item failed: %v", err)
			}
		}
	}()

	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return fmt.Errorf("must supply type")
	}

	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.ID != "" {
		indexService.Id(item.ID)
	}

	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}

	logger.Info("%+v", resp)
	return nil
}
