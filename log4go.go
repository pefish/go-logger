package go_logger

import (
	"fmt"
	"github.com/pefish/go-logger/log4go"
	"runtime/debug"
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

func (log4goInstance *Log4goClass) Init(prefix string, level string) {
	if prefix != `` {
		log4goInstance.prefix = fmt.Sprintf("[%s]: ", prefix)
	}
	sl := make(log4go.Logger)
	sl.AddFilter(`console`, log4goErrLevels[level], log4go.NewConsoleLogWriter())
	log4goInstance.logger = &sl
}

func (log4goInstance *Log4goClass) Close() {
	if log4goInstance.logger != nil {
		log4goInstance.logger.Close()
	}
}

func (log4goInstance *Log4goClass) Debug(args ...interface{}) {
	log4goInstance.logger.DebugFull("%s%s", log4goInstance.prefix, log4goInstance.FormatOutput(args...))
}

func (log4goInstance *Log4goClass) DebugF(format string, args ...interface{}) {
	log4goInstance.logger.DebugFull("%s%s", log4goInstance.prefix, fmt.Sprintf(format, args...))
}

func (log4goInstance *Log4goClass) Info(args ...interface{}) {
	log4goInstance.logger.InfoFull("%s%s", log4goInstance.prefix, log4goInstance.FormatOutput(args...))
}

func (log4goInstance *Log4goClass) InfoF(format string, args ...interface{}) {
	log4goInstance.logger.InfoFull("%s%s", log4goInstance.prefix, fmt.Sprintf(format, args...))
}

func (log4goInstance *Log4goClass) Warn(args ...interface{}) {
	log4goInstance.logger.WarnFull("%s%s", log4goInstance.prefix, log4goInstance.FormatOutput(args...))
}

func (log4goInstance *Log4goClass) WarnF(format string, args ...interface{}) {
	log4goInstance.logger.WarnFull("%s%s", log4goInstance.prefix, fmt.Sprintf(format, args...))
}

func (log4goInstance *Log4goClass) Error(args ...interface{}) {
	log4goInstance.logger.ErrorFull("%s%s", log4goInstance.prefix, log4goInstance.FormatOutput(args...))
}

func (log4goInstance *Log4goClass) ErrorF(format string, args ...interface{}) {
	log4goInstance.logger.ErrorFull("%s%s", log4goInstance.prefix, fmt.Sprintf(format, args...))
}

func (log4goInstance *Log4goClass) ErrorWithStack(args ...interface{}) {
	log4goInstance.logger.ErrorFull("%s%s\n%s", log4goInstance.prefix, log4goInstance.FormatOutput(args...), string(debug.Stack()))
}

func (log4goInstance *Log4goClass) ErrorWithStackF(format string, args ...interface{}) {
	log4goInstance.logger.ErrorFull("%s%s\n%s", log4goInstance.prefix, fmt.Sprintf(format, args...), string(debug.Stack()))
}
