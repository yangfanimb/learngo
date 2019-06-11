package parser

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestParseCityList(t *testing.T) {
	filename := "citylist.html"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read contents from %s failed: %v", filename, err)
	}

	result := ParseCityList(contents)

	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d request; but have %d", resultSize, len(result.Requests))
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}
}
