package main

import (
	"flag"
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"
	"learngo/crawler/scheduler"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	saver_client "learngo/crawler_distributed/persist/client"
	"learngo/crawler_distributed/rpcsupport"
	worker_client "learngo/crawler_distributed/worker/client"
	"net/rpc"
	"strings"
)

var logger = log.DLogger()

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err != nil {
			logger.Error("error connecting to %s: %v", h, err)
		} else {
			clients = append(clients, client)
			logger.Info("connect to %s", h)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}

var (
	itemSaverHost = flag.String("itemsaver_host", "", "itemsaver host")
	workerHosts   = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := saver_client.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	process, err := worker_client.CreateProcessor(pool)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      50,
		ItemChan:         itemChan,
		RequestProcessor: process,
	}

	e.Run(engine.Request{
		Url:    "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

}
