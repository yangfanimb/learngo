package engine

import (
	"learngo/crawler/helper/log"
)

var logger = log.DLogger()

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		for _, item := range parseResult.Items {
			logger.Info("Got item %v\n", item)
		}

		requests = append(requests, parseResult.Requests...)
	}

	logger.Info("finish all")
}
