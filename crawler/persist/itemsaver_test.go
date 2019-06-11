package persist

import (
	"context"
	"encoding/json"
	"learngo/crawler/engine"
	"learngo/crawler/model"
	"testing"

	"github.com/olivere/elastic"
)

func TestSave(t *testing.T) {
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

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.66.102:9200"), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	err = Save(client, "dating_test", expected)
	if err != nil {
		panic(err)
	}

	resp, err := client.Get().
		Index("dating_test").
		Type(expected.Type).
		Id(expected.ID).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(*(resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, err := model.FromJsonObj(actual.Payload)
	if err != nil {
		panic(err)
	}

	actual.Payload = actualProfile

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}
}
