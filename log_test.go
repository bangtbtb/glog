package glog

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var location, _ = filepath.Abs("./static")

// var location = ""
var config = &Config{
	Append:     true,
	BuildPath:  os.Getenv("GOPATH") + "/src/webapi/glogger",
	Location:   location,
	CodeLine:   true,
	IgnoreTags: []string{"ignore"},
}

func TestAll(t *testing.T) {
	assert.Equal(t, nil, config.NormalizeByEnv())
	config.InitResource("test-log")

	var manager = config.Logger
	manager.SetLevel(5)

	var log = manager.GetTagLog("TT")

	log.Empty()

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

	manager.Logger.Empty()
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
	iglog.Empty()

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
}
