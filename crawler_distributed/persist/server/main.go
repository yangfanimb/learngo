package main

import (
	"flag"
	"fmt"
	"learngo/crawler_distributed/config"
	"learngo/crawler_distributed/persist"
	"learngo/crawler_distributed/rpcsupport"
	"log"

	"github.com/olivere/elastic"
)


var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()

	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndx))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://192.168.66.102:9200"),
		elastic.SetSniff(false),
	)

	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
