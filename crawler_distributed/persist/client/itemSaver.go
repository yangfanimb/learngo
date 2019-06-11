package client

import (
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
)

var logger = log.DLogger()

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
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

			result := ""
			err := client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				logger.Error("save item failed: %v", err)
			}
		}
	}()
	return out, nil
}
