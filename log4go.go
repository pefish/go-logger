package go_logger

import (
	"fmt"
	"github.com/pefish/go-logger/log4go"
)

type Log4goClass struct {
	BaseLogger
	logger *log4go.Logger
	prefix string
}

var log4goErrLevels = map[string]log4go.Level{
	`debug`: log4go.DEBUG,
	`info`: log4go.INFO,
	`warn`: log4go.WARNING,
	`error`: log4go.ERROR,
}

func (this *Log4goClass) Init(prefix string, level string) {
	if prefix != `` {
		this.prefix = fmt.Sprintf("[%s]: ", prefix)
	}
	sl := make(log4go.Logger)
	sl.AddFilter(`console`, log4goErrLevels[level], log4go.NewConsoleLogWriter())
	this.logger = &sl
}

func (this *Log4goClass) Close() {
	if this.logger != nil {
		this.logger.Close()
	}
}

func (this *Log4goClass) Debug(args ...interface{}) {
	this.logger.DebugFull("%s%s", this.prefix, this.FormatOutput(args...))
}

func (this *Log4goClass) DebugF(format string, args ...interface{}) {
	this.logger.DebugFull("%s%s", this.prefix, fmt.Sprintf(format, args...))
}

func (this *Log4goClass) Info(args ...interface{}) {
	this.logger.InfoFull("%s%s", this.prefix, this.FormatOutput(args...))
}

func (this *Log4goClass) InfoF(format string, args ...interface{}) {
	this.logger.InfoFull("%s%s", this.prefix, fmt.Sprintf(format, args...))
}

func (this *Log4goClass) Warn(args ...interface{}) {
	this.logger.WarnFull("%s%s", this.prefix, this.FormatOutput(args...))
}

func (this *Log4goClass) WarnF(format string, args ...interface{}) {
	this.logger.WarnFull("%s%s", this.prefix, fmt.Sprintf(format, args...))
}

func (this *Log4goClass) Error(args ...interface{}) {
	this.logger.ErrorFull("%s%s", this.prefix, this.FormatOutput(args...))
}

func (this *Log4goClass) ErrorF(format string, args ...interface{}) {
	this.logger.ErrorFull("%s%s", this.prefix, fmt.Sprintf(format, args...))
}
