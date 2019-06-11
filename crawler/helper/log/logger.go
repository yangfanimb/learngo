package log

import (
	"os"

	"github.com/astaxie/beego/logs"
)

func DLogger() *logs.BeeLogger {
	return logs.GetBeeLogger()
}

func init() {
	//logs.EnableFuncCallDepth(true)
	//logs.SetLogFuncCallDepth(3)
	os.Remove("crawler.log")
	logs.SetLogger(logs.AdapterFile, `{"filename":"crawler.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.SetLogger(logs.AdapterConsole, "")
}
