package glog

import (
	"strconv"
	"time"
)

// SERVICE_NAME ||	Type || Time || Content
// SERVICE_NAME ||	Type || Time || [tag] || Content
// TEST || DBG || 19:3:26 || /home/bangnl/path.go:45:Accept ||  [TAG] Done OnJoin. Start receive loop
const (
	loggerTraceHead   = "TRC"
	loggerDebugHead   = "DBG"
	loggerInfoHead    = "INF"
	loggerErrorHead   = "ERR"
	loggerWarningHead = "WRN"
	loggerFatalHead   = "FTL"
	loggerAll         = "LOGGER "
	bulkhead          = "||"
	bulkheadSpace     = " || "
)

// Log level
const (
	LevelTrace = 10
	LevelDebug = 11
	LevelInfo  = 12
	LevelWarn  = 13
	LevelError = 14
	LevelFatal = 15
)

//AppendLog :
var AppendLog = true

func strTime() string {
	now := time.Now()
	return strconv.Itoa(now.Year()%2000) +
		":" + strconv.Itoa(int(now.Month())) +
		":" + strconv.Itoa(now.Day()) //+ "▶"
}
