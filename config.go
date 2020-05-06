package glog

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
)

//Config : config of this package
type Config struct {
	Level      int         `json:"level"`
	ShowLine   bool        `json:"showLine"`
	ShowStd    bool        `json:"showStd"`
	Append     bool        `json:"append"`
	AppName    string      `json:"appName"`
	BuildPath  string      `json:"buildPath" xml:"buildPath"`
	Location   string      `json:"location" xml:"location"`
	IgnoreTags []string    `json:"igTags" xml:"igTags"`
	Logger     *TagManager `json:"-"`
}

//NormalizeByEnv :
func (config *Config) NormalizeByEnv() error {

	if os.Getenv("LOG_APP_NAME") != "" {
		config.AppName = os.Getenv("LOG_APP_NAME")
	}

	if os.Getenv("LOG_LEVEL") != "" {
		level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if nil != err {
			return err
		}
		config.Level = level
	}

	if os.Getenv("LOG_LOCATION") != "" {
		var location, err = filepath.Abs(os.Getenv("LOG_LOCATION"))
		if nil != err {
			return err
		}
		config.Location = location
	}

	return nil
}

//InitResource :
func (config *Config) InitResource() *TagManager {
	logex, err := NewTagManager(config)
	if nil == err {
		logex.GetTagLog("SHARED").Info("Init database and logger success")
	}
	config.Logger = logex
	if nil != err {
		panic(errors.New("Init glogger error: " + err.Error()))
	}
	return logex
}
