package go_logger

import (
	"errors"
	"fmt"
	"github.com/pefish/log4go"
	"os"
)

type Log4goClass struct {
	BaseLogger
	logger *log4go.Logger
}

func (this *Log4goClass) Init(name string, debug bool) {
	sl := make(log4go.Logger)
	level := log4go.INFO
	if debug {
		level = log4go.DEBUG
		sl.AddFilter(`console`, level, log4go.NewConsoleLogWriter())
	}

	logfile := os.Getenv(`GO_LOG`)
	if logfile != `` {
		logWriter := log4go.NewFileLogWriter(logfile + fmt.Sprintf(`/%s.log`, name), true, true)
		if logWriter == nil {
			panic(errors.New(`GO_LOG config error`))
		}
		sl.AddFilter("file", level, logWriter)
	}
	this.logger = &sl
}

func (this *Log4goClass) Close() {
	if this.logger != nil {
		this.logger.Close()
	}
}


func (this *Log4goClass) Debug(args ...interface{}) {
	this.logger.DebugFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) Info(args ...interface{}) {
	this.logger.InfoFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) Warn(args ...interface{}) {
	this.logger.WarnFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) Error(args ...interface{}) {
	this.logger.ErrorFull(`%s`, this.FormatOutput(args...))
}

