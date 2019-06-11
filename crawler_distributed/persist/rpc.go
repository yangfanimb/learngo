package persist

import (
	"learngo/crawler/engine"
	"learngo/crawler/helper/log"
	"learngo/crawler/persist"

	"github.com/olivere/elastic"
)

var logger = log.DLogger()

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	if err != nil {
		logger.Error("Error saving item %v: %v", item, err)
		return err
	}

	*result = "ok"
	logger.Info("Item %v saved", item)
	return nil
}
