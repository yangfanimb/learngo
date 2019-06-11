package config

const (
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"

	// Service port
	//ItemSaverPort = 1234
	//WorkerPort0   = 9000

	// ElasticSearch
	ElasticIndx = "dating_profile"

	// RPC Endpoints
	ItemSaverRpc    = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	Qps = 20
)
