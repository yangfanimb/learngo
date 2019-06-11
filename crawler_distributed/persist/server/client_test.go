package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestClientSaver(t *testing.T) {
	const host = ":1234"

	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// call save
	item := engine.Item{
		URL:  "http://album.zhenai.com/u/1320662004",
		Type: "zhenai",
		ID:   "1320662004",
		Payload: model.Profile{
			Name:       "非诚勿扰",
			Gender:     "男士",
			Age:        29,
			Height:     180,
			Weight:     66,
			Income:     "1.2-2万",
			Marriage:   "未婚",
			Education:  "高中及以下",
			Occupation: "",
			WorkArea:   "阿坝茂县",
			Hukou:      "四川阿坝",
			Xinzuo:     "天秤座(09.23-10.22)",
			House:      "已购房",
			Car:        "已买车",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result %s: error: %v", result, err)
	}
}
