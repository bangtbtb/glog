package glog

import (
	"errors"
	"io"
	"os"
	"strconv"
	"time"
)

func newFileWriter(dir, filePrefix string) (io.ReadWriteCloser, error) {
	if dir == "" {
		return nil, errors.New("log directory is not valid")
	}
	os.MkdirAll(dir, 0755)

	t := time.Now()
	ts := strconv.Itoa(t.Year()) + "-" + strconv.Itoa(int(t.Month())) + "-" + strconv.Itoa(int(t.Day()))
	fileName := filePrefix + "-" + ts + "-" + strconv.FormatInt(t.Unix(), 10) + ".log"

	if len(dir) > 1 && dir[len(dir)-1] != '/' {
		dir += "/"
	}

	fileName = dir + fileName
	return os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
}
