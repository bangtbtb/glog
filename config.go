package glog

import "errors"

//Config : config of this package
type Config struct {
	CodeLine   bool        `json:"codeLine"`
	Append     bool        `json:"append"`
	BuildPath  string      `json:"buildPath" xml:"buildPath"`
	Location   string      `json:"location" xml:"location"`
	IgnoreTags []string    `json:"igTags" xml:"igTags"`
	Logger     *LogManager `json:"-"`
}

//NormalizeByEnv :
func (config *Config) NormalizeByEnv() error {
	return nil
}

//InitResource :
func (config *Config) InitResource(prjName string) *LogManager {
	logex, err := NewLogManager(config.IgnoreTags, config.BuildPath, prjName, config.Location, config.CodeLine)
	if nil == err {
		logex.GetTagLog("SHARED").Info("Init database and logger success")
	}
	config.Logger = logex
	if nil != err {
		panic(errors.New("Init glogger error: " + err.Error()))
	}
	return logex
}
