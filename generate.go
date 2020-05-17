package glog

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

//test-log || FTL || 19:9:13 || /log_test.go:41 -> /usr/local/etc/go-v1.12.7/src/testing/testing.go:868  ||  [TT] fatal 1

var minLen = strings.Repeat(" ", 60)

type logGenerate struct {
	codeLine   bool
	trimPrefix string
	logFormat  string
}

func newLogGenerate(printCode bool, header string, trimPrefix string) *logGenerate {
	var gen = &logGenerate{codeLine: printCode, trimPrefix: trimPrefix}
	if gen.codeLine {
		gen.logFormat = header + bulkheadSpace + "%s" + bulkheadSpace + "%s" + bulkheadSpace + "%s" + bulkheadSpace
	} else {
		gen.logFormat = header + bulkheadSpace + "%s" + bulkheadSpace + "%s" + bulkheadSpace
	}
	return gen
}

func (gen *logGenerate) genLogPrefix(skip, depth int, typeHead string) string {
	if depth < 1 {
		depth = 1
	}
	var s string
	var msg string
	if gen.codeLine {
		var pc = make([]uintptr, depth)
		runtime.Callers(skip+1, pc)
		for i := 0; i < len(pc); i++ {
			if pc[i] == 0 {
				break
			}
			file, line := runtime.FuncForPC(pc[i]).FileLine(pc[i])
			if len(file) > 8 && file[len(file)-8:] == "panic.go" {
				continue
			}
			gen.splitFilePath(&file)
			s = fmt.Sprintf("%s:%d", file, line)
			msg += s + " -> "
		}
		s = msg[:len(msg)-3]
		return fmt.Sprintf(gen.logFormat, typeHead, strTime(), s)
	}
	return fmt.Sprintf(gen.logFormat, typeHead, strTime())
}

func (gen *logGenerate) splitFilePath(path *string) {
	s, err := filepath.Abs(*path)
	if nil == err {
		index := strings.Index(s, gen.trimPrefix)
		if index != -1 {
			*path = (s)[len(gen.trimPrefix):]
		}
	}
}
