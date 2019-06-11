package engine

type ParserFunc func([]byte) ParseResult
type Parser interface {
	Parse(contents []byte) ParseResult
	Serialize() (name string, args interface{})
}

type Request struct {
	Url    string
	Parser Parser
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	URL     string
	Type    string
	ID      string
	Payload interface{}
}

type NilParser struct{}

func (NilParser) Parse(_ []byte) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}

func (f *FuncParser) Parse(contents []byte) ParseResult {
	return f.parser(contents)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}
