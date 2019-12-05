package glog

//IBaseLog :
type IBaseLog interface {
	ISetting
	ILogger
}

//ISetting :
type ISetting interface {
	SetLevel(level int)
	SetEventSignal(c chan string)
	Flush()
}

//ILogger : Logger interface
type ILogger interface {
	//EmptyBlock
	Empty()

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
