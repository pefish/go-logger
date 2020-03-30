package go_logger

import (
	"errors"
	"fmt"
	"github.com/pefish/go-logger/log4go"
	"os"
)

type Log4goClass struct {
	BaseLogger
	logger *log4go.Logger
}

var log4goErrLevels = map[string]log4go.Level{
	`debug`: log4go.DEBUG,
	`info`: log4go.INFO,
	`warn`: log4go.WARNING,
	`error`: log4go.ERROR,
}

func (this *Log4goClass) Init(name string, level string) {
	sl := make(log4go.Logger)
	sl.AddFilter(`console`, log4goErrLevels[level], log4go.NewConsoleLogWriter(), name)
	logfile := os.Getenv(`GO_LOG`)
	if logfile != `` {
		logWriter := log4go.NewFileLogWriter(logfile+fmt.Sprintf(`/%s.log`, name), true, true)
		if logWriter == nil {
			panic(errors.New(`GO_LOG config error`))
		}
		sl.AddFilter("file", log4goErrLevels[level], logWriter, name)
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

func (this *Log4goClass) DebugF(format string, args ...interface{}) {
	this.logger.DebugFull(format, args...)
}

func (this *Log4goClass) Info(args ...interface{}) {
	this.logger.InfoFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) InfoF(format string, args ...interface{}) {
	this.logger.InfoFull(format, args...)
}

func (this *Log4goClass) Warn(args ...interface{}) {
	this.logger.WarnFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) WarnF(format string, args ...interface{}) {
	this.logger.WarnFull(format, args...)
}

func (this *Log4goClass) Error(args ...interface{}) {
	this.logger.ErrorFull(`%s`, this.FormatOutput(args...))
}

func (this *Log4goClass) ErrorF(format string, args ...interface{}) {
	this.logger.ErrorFull(format, args...)
}
