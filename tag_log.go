package glog

type tagLog struct {
	Logger ILogger
	Tag    string
}

func newTagLog(tag string, il ILogger) ILogger {
	var res = &tagLog{Tag: " [" + tag + "] ", Logger: il}
	return res
}

//Empty :
func (tl *tagLog) Empty() {
	tl.Logger.Empty()
}

func (tl *tagLog) Trace(format string, args ...interface{}) {
	tl.Logger.TraceStack(1, 1, tl.Tag+format, args...)
}

func (tl *tagLog) TraceStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.TraceStack(skip+1, depth, tl.Tag+format, args...)
}

//Debug :
func (tl *tagLog) Debug(format string, args ...interface{}) {
	tl.Logger.DebugStack(1, 1, tl.Tag+format, args...)
}

//DebugStack :
func (tl *tagLog) DebugStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.DebugStack(skip+1, depth, tl.Tag+format, args...)
}

//Infoln :
func (tl *tagLog) Info(format string, args ...interface{}) {
	tl.Logger.InfoStack(1, 1, tl.Tag+format, args...)
}

func (tl *tagLog) InfoStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.InfoStack(skip+1, depth, tl.Tag+format, args...)
}

func (tl *tagLog) Error(format string, args ...interface{}) {
	tl.Logger.ErrorStack(1, 1, tl.Tag+format, args...)
}

func (tl *tagLog) ErrorStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.ErrorStack(skip+1, depth, tl.Tag+format, args...)
}

func (tl *tagLog) Warn(format string, args ...interface{}) {
	tl.Logger.WarnStack(1, 1, tl.Tag+format, args...)
}

func (tl *tagLog) WarnStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.WarnStack(skip+1, depth, tl.Tag+format, args...)
}

func (tl *tagLog) Fatal(format string, args ...interface{}) {
	tl.Logger.FatalStack(1, 1, tl.Tag+format, args...)
}

func (tl *tagLog) FatalStack(skip, depth int, format string, args ...interface{}) {
	tl.Logger.FatalStack(skip+1, depth, tl.Tag+format, args...)
}
