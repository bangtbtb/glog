package glog

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var location, _ = filepath.Abs("./static")

// var location = ""
var config = &Config{
	Append:     true,
	AppName:    "test-log",
	BuildPath:  os.Getenv("GOPATH") + "/src/github.com/bangtbtb/glog",
	Location:   location,
	ShowLine:   true,
	ShowStd:    true,
	IgnoreTags: []string{"ignore"},
}

func TestAll(t *testing.T) {
	os.Setenv("LOG_APP_NAME", "glogger-test")
	os.Setenv("LOG_LEVEL", "0")
	os.Setenv("LOG_LOCATION", "./static/log")
	CheckError("Init by environment", config.NormalizeByEnv())

	var manager = config.InitResource()
	var pipe = make(chan *LogMsg, 256)
	var ctx, cancel = context.WithCancel(context.TODO())

	go listenChannel(ctx, pipe)
	manager.SetLevel(5)
	manager.Logger.SetEventSignal(pipe)

	var log = manager.GetTagLog("TT")

	log.Trace("Trace at %d", 1)
	log.TraceStack(0, 2, "Trace stack at %s", "abc")

	log.Debug("debug at %d", 1)
	log.DebugStack(0, 2, "debug stack at %s", "abc")

	log.Info("info at %d", 1)
	log.InfoStack(0, 2, "info stack at %s", "abc")

	log.Warn("warn at %d", 1)
	log.WarnStack(0, 2, "warn stack at %s", "abc")

	log.Error("err at %d", 1)
	log.ErrorStack(0, 2, "error stack at %s", "abc")

	log.Fatal("fatal %d", 1)
	log.FatalStack(0, 2, "fatal %d", 1)

	manager.Logger.Trace("Trace at %d", 1)
	manager.Logger.TraceStack(0, 2, "Trace stack at %s", "abc")
	manager.Logger.Debug("debug at %d", 1)
	manager.Logger.DebugStack(0, 2, "debug stack at %s", "abc")
	manager.Logger.Info("info at %d", 1)
	manager.Logger.InfoStack(0, 2, "info stack at %s", "abc")
	manager.Logger.Warn("warn at %d", 1)
	manager.Logger.WarnStack(0, 2, "warn stack at %s", "abc")
	manager.Logger.Error("err at %d", 1)
	manager.Logger.ErrorStack(0, 2, "error stack at %s", "abc")
	manager.Logger.Fatal("fatal %d", 1)
	manager.Logger.FatalStack(0, 2, "fatal %d", 1)

	var iglog = manager.GetTagLog("ignore")

	iglog.Trace("Trace at %d", 1)
	iglog.TraceStack(0, 2, "Trace stack at %s", "abc")

	iglog.Debug("debug at %d", 1)
	iglog.DebugStack(0, 2, "debug stack at %s", "abc")

	iglog.Info("info at %d", 1)
	iglog.InfoStack(0, 2, "info stack at %s", "abc")

	iglog.Warn("warn at %d", 1)
	iglog.WarnStack(0, 2, "warn stack at %s", "abc")

	iglog.Error("err at %d", 1)
	iglog.ErrorStack(0, 2, "error stack at %s", "abc")

	iglog.Fatal("fatal %d", 1)
	iglog.FatalStack(0, 2, "fatal %d", 1)

	manager.Logger.Flush()
	cancel()
}

func listenChannel(ctx context.Context, pipe chan *LogMsg) {
	var done = false
	for {
		select {
		case <-ctx.Done():
			done = true
			break
		case msg := <-pipe:
			fmt.Printf("Receive log: %+v\n", msg)
			break
		}
		if done {
			break
		}
	}
}

//CheckError :
func CheckError(msg string, err error) {
	if nil != err {
		log.Fatalln(msg+": ", err)
	}
}
