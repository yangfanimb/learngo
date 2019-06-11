package parser

import (
	"io/ioutil"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		t.Fatalf("read profile data file failed: %v", err)
	}

	item := engine.Item{
		URL:  "http://album.zhenai.com/u/1320662004",
		Type: "zhenai",
		Payload: model.Profile{
			Gender: "男士",
		},
	}
	result := ParseProfile(contents, item)

	if len(result.Items) != 1 {
		t.Errorf("items should contain 1 element; but was %v", result.Items)
	}

	item = result.Items[0]
	expected := engine.Item{
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

	if item != expected {
		t.Errorf("expected %v; but was %v", expected, item)
	}
}
