package glog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gookit/color"
)

//fileLogger : Base Logger
type fileLogger struct {
	sync.Mutex
	level     int
	appName   string
	showstd   bool
	gen       *logGenerate
	writer    io.ReadWriteCloser
	signalOut chan *LogMsg
}

//OptionFileLogger : use for create new file logger
type OptionFileLogger struct {
	ShowStd    bool               // Write log to stdout
	ShowLine   bool               // Show code line
	TrimPrefix string             // Trim prefix of code line when run
	Name       string             // Name of log
	LogDir     string             // If logdir or writer was not set
	Writer     io.ReadWriteCloser //
}

//NewFileLogger : Create file logger
func NewFileLogger(option OptionFileLogger) (ILogWriter, error) {
	var err error
	var writer io.ReadWriteCloser
	if option.LogDir != "" {
		writer, err = newFileWriter(option.LogDir, option.Name)
		if nil != err {
			return nil, err
		}
	}

	var res = &fileLogger{appName: option.Name}
	res.showstd = option.ShowStd
	res.gen = newLogGenerate(option.ShowLine, option.Name, option.TrimPrefix)
	res.writer = writer
	return res, nil
}

//NewStdLog :
func NewStdLog(trimPrefix, name string, codeLine bool) ILogWriter {
	var res = &fileLogger{}
	res.gen = newLogGenerate(codeLine, name, trimPrefix)
	res.writer = os.Stdout
	return res
}

func (bl *fileLogger) SetLevel(level int) {
	bl.level = level
}

func (bl *fileLogger) write(msg string, level int, writer io.Writer) {
	bl.Lock()
	defer bl.Unlock()
	if writer != nil {
		writer.Write([]byte(msg + "\r\n"))
	}
	if bl.showstd {
		os.Stdout.Write([]byte(msg + "\r\n"))
	}
	if nil != bl.signalOut {
		bl.signalOut <- &LogMsg{Level: level, Time: time.Now().Unix(), Msg: msg, AppName: bl.appName}
	}
}

func (bl *fileLogger) SetEventSignal(c chan *LogMsg) {
	bl.signalOut = c
}

func (bl *fileLogger) Flush() {
	if bl.writer != nil {
		bl.writer.Close()
	}
}

func (bl *fileLogger) Trace(format string, args ...interface{}) {
	if bl.level <= LevelTrace {
		p := bl.gen.genLogPrefix(2, 1, loggerTraceHead) + fmt.Sprintf(format, args...)

		bl.write(p, LevelTrace, bl.writer)
		color.Gray.Println(p)
	}
}

func (bl *fileLogger) TraceStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelTrace {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerTraceHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelTrace, bl.writer)
		color.Gray.Println(p)
	}
}

func (bl *fileLogger) Debug(format string, args ...interface{}) {
	if bl.level <= LevelDebug {
		p := bl.gen.genLogPrefix(2, 1, loggerDebugHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelDebug, bl.writer)
		color.Cyan.Println(p)
	}
}

func (bl *fileLogger) DebugStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelDebug {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerDebugHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelDebug, bl.writer)
		color.Cyan.Println(p)
	}
}

func (bl *fileLogger) Info(format string, args ...interface{}) {
	if bl.level <= LevelInfo {
		p := bl.gen.genLogPrefix(2, 1, loggerInfoHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelInfo, bl.writer)
		color.Green.Println(p)
	}
}

func (bl *fileLogger) InfoStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelInfo {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerInfoHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelInfo, bl.writer)
		color.Green.Println(p)
	}
}

func (bl *fileLogger) Warn(format string, args ...interface{}) {
	if bl.level <= LevelWarn {
		p := bl.gen.genLogPrefix(2, 1, loggerWarningHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelWarn, bl.writer)
		color.Yellow.Println(p)
	}

}

func (bl *fileLogger) WarnStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelWarn {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerWarningHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelWarn, bl.writer)
		color.Yellow.Println(p)
	}
}

func (bl *fileLogger) Error(format string, args ...interface{}) {
	if bl.level <= LevelError {
		p := bl.gen.genLogPrefix(2, 1, loggerErrorHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelError, bl.writer)
		color.Error.Println(p)
	}
}

func (bl *fileLogger) ErrorStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelError {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerErrorHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelError, bl.writer)
		color.Error.Println(p)
	}
}

func (bl *fileLogger) Fatal(format string, args ...interface{}) {
	if bl.level <= LevelFatal {
		p := bl.gen.genLogPrefix(2, 1, loggerFatalHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelFatal, bl.writer)
		color.Red.Println(p)
	}
}

func (bl *fileLogger) FatalStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelFatal {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerFatalHead) + fmt.Sprintf(format, args...)
		bl.write(p, LevelFatal, bl.writer)
		color.Red.Println(p)
	}
}
