package glog

//ILogWriter :
type ILogWriter interface {
	ILogger
	Flush()
	SetLevel(level int)
	SetEventSignal(c chan *LogMsg)
}

//ILogger : Logger interface
type ILogger interface {
	Trace(format string, args ...interface{})
	TraceStack(skip, depth int, format string, args ...interface{})

	Debug(format string, args ...interface{})
	DebugStack(skip, depth int, format string, args ...interface{})

	Info(format string, args ...interface{})
	InfoStack(skip, depth int, format string, args ...interface{})

	Error(format string, args ...interface{})
	ErrorStack(skip, depth int, format string, args ...interface{})

	Warn(format string, args ...interface{})
	WarnStack(skip, depth int, format string, args ...interface{})

	Fatal(format string, args ...interface{})
	FatalStack(skip, depth int, format string, args ...interface{})
}

//LogMsg : log struct
type LogMsg struct {
	ID      string `bson:"_id" json:"id"`                              //
	Level   int    `bson:"level,omitempty" json:"level,omitempty"`     //
	Time    int64  `bson:"time,omitempty" json:"time,omitempty"`       //
	Msg     string `bson:"msg,omitempty" json:"msg,omitempty"`         //
	AppName string `bson:"appName,omitempty" json:"appName,omitempty"` // App name
}
