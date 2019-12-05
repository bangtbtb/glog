package glog

//LogManager : Logger manager
type LogManager struct {
	Config  *Config
	Logger  IBaseLog
	Ignores []string
	TagLogs map[string]ILogger
}

//NewLogManager :
func NewLogManager(igs []string, trimPrefix, name, logdir string, codeLine bool) (*LogManager, error) {
	var manager = &LogManager{Ignores: igs}
	if nil == igs {
		manager.Ignores = make([]string, 0)
	}
	var err error
	manager.Logger, err = NewFileLogger(trimPrefix, name, logdir, codeLine)
	if nil != err {
		return nil, err
	}
	manager.Logger.SetLevel(0)
	manager.TagLogs = make(map[string]ILogger)
	return manager, nil
}

//GetTagLog :
func (tle *LogManager) GetTagLog(tag string) ILogger {
	ret, ok := tle.TagLogs[tag]
	if ok {
		return ret
	}
	var il ILogger
	if tle.isIgnore(tag) {
		il = NewIgnoreLog()
	} else {
		il = newTagLog(tag, tle.Logger)
	}
	tle.TagLogs[tag] = il
	return il
}

func (tle *LogManager) isIgnore(tag string) bool {
	for _, s := range tle.Ignores {
		if s == tag {
			return true
		}
	}
	return false
}

//SetLevel :
func (tle *LogManager) SetLevel(level int) {
	if nil != tle.Logger {
		tle.Logger.SetLevel(level)
	}
}
