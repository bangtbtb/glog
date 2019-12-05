package glog

//NoneLog :
type ignoreLog struct {
}

//NewIgnoreLog :
func NewIgnoreLog() ILogger {
	var res = &ignoreLog{}
	return res
}

func (nl *ignoreLog) Empty() {
}

func (nl *ignoreLog) Trace(format string, args ...interface{}) {}

func (nl *ignoreLog) TraceStack(skip, depth int, format string, args ...interface{}) {}

func (nl *ignoreLog) Debug(format string, args ...interface{}) {}

func (nl *ignoreLog) DebugStack(skip, depth int, format string, args ...interface{}) {}

func (nl *ignoreLog) Info(format string, args ...interface{}) {}

func (nl *ignoreLog) InfoStack(skip, depth int, format string, args ...interface{}) {}

func (nl *ignoreLog) Error(format string, args ...interface{}) {}

func (nl *ignoreLog) ErrorStack(skip, depth int, format string, args ...interface{}) {}

func (nl *ignoreLog) Warn(format string, args ...interface{}) {}

func (nl *ignoreLog) WarnStack(skip, depth int, format string, args ...interface{}) {}

func (nl *ignoreLog) Fatal(format string, args ...interface{}) {}

func (nl *ignoreLog) FatalStack(skip, depth int, format string, args ...interface{}) {}
