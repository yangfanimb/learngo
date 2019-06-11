package worker

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"
	"learngo/crawler/model"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
)

var logger = log.DLogger()

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Requests []Request
	Items    []engine.Item
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(r engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}

	return result
}

func DeserializeRequest(r Request) (engine.Request, error) {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			logger.Error("err deserializing request: %v err: %v", req, err)
			continue
		}

		result.Requests = append(result.Requests, engineReq)
	}

	return result
}

func deserializeParser(p SerializedParser) (engine.Parser, error) {
	switch p.Name {
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseProfile:
		var item engine.Item
		if m, ok := p.Args.(map[string]interface{}); ok{
			item.URL = m["URL"].(string)
			item.Type = m["Type"].(string)
			item.ID = m["ID"].(string)

			if mm, ok := m["Payload"].(map[string]interface{}); ok {
				var profile model.Profile
				profile.Gender = mm["Gender"].(string)

				item.Payload = profile
			}

			return parser.NewProfileParser(item), nil
		}

		return nil, fmt.Errorf("invalid arg: %v", p.Args)
	case config.NilParser:
		return engine.NilParser{}, nil
	default:
		return nil, fmt.Errorf("unknown parser name: %s", p.Name)
	}
}
