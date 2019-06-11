package engine

import (
	"learngo/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	logger.Info("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		logger.Warn("Fetcher:error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body), nil
}
