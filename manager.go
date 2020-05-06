package glog

//TagManager : Logger manager
type TagManager struct {
	Config  *Config
	Logger  ILogWriter
	Ignores []string
	TagLogs map[string]ILogger
}

//NewTagManager :
func NewTagManager(config *Config) (*TagManager, error) {
	var manager = &TagManager{Ignores: config.IgnoreTags}
	if nil == manager.Ignores {
		manager.Ignores = make([]string, 0)
	}
	var err error
	manager.Logger, err = NewFileLogger(OptionFileLogger{
		ShowStd:    config.ShowStd,
		ShowLine:   config.ShowLine,
		TrimPrefix: config.BuildPath,
		Name:       config.AppName,
		LogDir:     config.Location,
	})
	if nil != err {
		return nil, err
	}
	manager.Logger.SetLevel(0)
	manager.TagLogs = make(map[string]ILogger)
	return manager, nil
}

//GetTagLog :
func (tle *TagManager) GetTagLog(tag string) ILogger {
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

func (tle *TagManager) isIgnore(tag string) bool {
	for _, s := range tle.Ignores {
		if s == tag {
			return true
		}
	}
	return false
}

//SetLevel :
func (tle *TagManager) SetLevel(level int) {
	if nil != tle.Logger {
		tle.Logger.SetLevel(level)
	}
}
