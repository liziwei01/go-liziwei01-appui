package logs

import (
	"github.com/astaxie/beego/logs"
)

func InitLoggers() {
	log := logs.NewLogger(10000)
	// log.SetLogger("console", "")
	log.SetLogger("file", `{"filename":"./log/test.log"}`)
}
