package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"time"
)

const (
	DEFAULT_PREFIX = "log"
	DEFAULT_DIR    = "logs"
)

type DefaultLog struct {
	prefix string
	dir    string
	file   *os.File
	ctime  time.Time
}

func init() {
	defaultLog := NewDailyWriteLog(DEFAULT_PREFIX, DEFAULT_DIR)
	logrus.SetOutput(defaultLog)
	logrus.SetLevel(logrus.DebugLevel)
}

func logTrace() *logrus.Entry {
	_, file, line, _ := runtime.Caller(2)
	trace := logrus.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	})
	return trace
}

func Info(data ...interface{}) {
	trace := logTrace()
	trace.Info(data)
}

func (dl *DefaultLog) Write(p []byte) (n int, err error) {
	dl.checkTime()
	n, err = dl.file.Write(p)
	return n, err
}

func (dl *DefaultLog) checkTime() {
	if !time.Now().Before(dl.ctime) {
		dl.ctime = time.Now().AddDate(0, 0, 1)
		dl.initFile()
	}
}

func (dl *DefaultLog) initFile() {
	if dl.file != nil {
		dl.file.Close()
	}
	if dl.dir == "" {
		dl.dir, _ = os.Getwd()
	}
	//if dl.dir != "" {
	//	dl.dir = strings.Trim(dl.dir, "/")
	//	logPath = fmt.Sprintf("%s/%s", logPath, dl.dir)
	//}
	if _, err := os.Stat(dl.dir); err != nil {
		err := os.MkdirAll(dl.dir, 0777)
		if err != nil {
			panic(err)
		}
	}

	timeFormat := dl.ctime.Format("2006-01-02")
	if dl.prefix != "" {
		timeFormat = "-" + timeFormat
	}
	filename := fmt.Sprintf("%s/%s%s.log", dl.dir, dl.prefix, timeFormat)
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	dl.file = file
}

func NewDailyWriteLog(prefix string, dir ...string) *DefaultLog {
	instance := new(DefaultLog)
	instance.prefix = prefix
	instance.ctime = time.Now().AddDate(0, 0, 1)
	if len(dir) > 0 {
		instance.dir = dir[0]
	}
	instance.initFile()

	return instance
}
