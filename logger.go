package glog

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/gookit/color"
)

//fileLogger : Base Logger
type fileLogger struct {
	sync.Mutex
	level       int
	gen         *logGenerate
	Writer      io.ReadWriteCloser
	EventSignal chan string
}

//NewFileLogger : Create file logger
func NewFileLogger(trimPrefix, name, logDir string, codeLine bool) (IBaseLog, error) {
	var err error
	var writer io.ReadWriteCloser
	if logDir != "" {
		writer, err = newFileWriter(logDir, name)
		if nil != err {
			return nil, err
		}
	} else {
		writer = os.Stdout
	}

	var res = &fileLogger{}
	res.gen = newLogGenerate(codeLine, name, trimPrefix)
	res.Writer = writer
	return res, nil
}

func (bl *fileLogger) SetLevel(level int) {
	bl.level = level
}

func (bl *fileLogger) write(msg string, writer io.Writer) {
	bl.Lock()
	defer bl.Unlock()
	if writer != nil {
		writer.Write([]byte(msg + "\r\n"))
	}

	if nil != bl.EventSignal {
		bl.EventSignal <- msg
	}
}

func (bl *fileLogger) SetEventSignal(c chan string) {
	bl.EventSignal = c
}

func (bl *fileLogger) Flush() {
	if bl.Writer != nil {
		bl.Writer.Close()
	}
}

func (bl *fileLogger) Empty() {
	bl.write("\n\r\n\r", bl.Writer)
}

func (bl *fileLogger) Trace(format string, args ...interface{}) {
	if bl.level <= LevelTrace {
		// p := bl.headerStackLog(2, 1, loggerTraceHead) + fmt.Sprintf(format, args...)
		p := bl.gen.genLogPrefix(2, 1, loggerTraceHead) + fmt.Sprintf(format, args...)

		bl.write(p, bl.Writer)
		color.Gray.Println(p)
	}
}

func (bl *fileLogger) TraceStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelTrace {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerTraceHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Gray.Println(p)
	}
}

func (bl *fileLogger) Debug(format string, args ...interface{}) {
	if bl.level <= LevelDebug {
		p := bl.gen.genLogPrefix(2, 1, loggerDebugHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Cyan.Println(p)
	}
}

func (bl *fileLogger) DebugStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelDebug {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerDebugHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Cyan.Println(p)
	}
}

func (bl *fileLogger) Info(format string, args ...interface{}) {
	if bl.level <= LevelInfo {
		p := bl.gen.genLogPrefix(2, 1, loggerInfoHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Green.Println(p)
	}
}

func (bl *fileLogger) InfoStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelInfo {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerInfoHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Green.Println(p)
	}
}

func (bl *fileLogger) Warn(format string, args ...interface{}) {
	if bl.level <= LevelWarn {
		p := bl.gen.genLogPrefix(2, 1, loggerWarningHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Yellow.Println(p)
	}

}

func (bl *fileLogger) WarnStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelWarn {
		p := bl.gen.genLogPrefix(2+1, depth, loggerWarningHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Yellow.Println(p)
	}
}

func (bl *fileLogger) Error(format string, args ...interface{}) {
	if bl.level <= LevelError {
		p := bl.gen.genLogPrefix(2, 1, loggerErrorHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Error.Println(p)
	}
}

func (bl *fileLogger) ErrorStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelError {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerErrorHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Error.Println(p)
	}
}

func (bl *fileLogger) Fatal(format string, args ...interface{}) {
	if bl.level <= LevelFatal {
		p := bl.gen.genLogPrefix(2, 1, loggerFatalHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Red.Println(p)
	}
}

func (bl *fileLogger) FatalStack(skip, depth int, format string, args ...interface{}) {
	if bl.level <= LevelFatal {
		p := bl.gen.genLogPrefix(2+skip, depth, loggerFatalHead) + fmt.Sprintf(format, args...)
		bl.write(p, bl.Writer)
		color.Red.Println(p)
	}
}
