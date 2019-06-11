package main

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"

	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://album.zhenai.com/u/1320662004",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: engine.Item{
				URL:"http://album.zhenai.com/u/1320662004",
				Type:"zhenai",
				ID:"1320662004",
				Payload: model.Profile{
					Gender:"男士",
				},
			},
		},
	}

	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("rpc call failed: %v", err)
	} else {
		fmt.Println(result)
	}
}
