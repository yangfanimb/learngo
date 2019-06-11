package client

import (
	"learngo/crawler/engine"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) (engine.Processor, error) {
	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResult worker.ParseResult

		c := <-clientChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResult)
		if err != nil {
			return engine.ParseResult{}, err
		}

		return worker.DeserializeResult(sResult), nil
	}, nil
}
