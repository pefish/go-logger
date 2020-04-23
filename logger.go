package go_logger

import (
	"github.com/pefish/go-interface-logger"
)

type LoggerClass struct {
	logger go_interface_logger.InterfaceLogger
}


var Logger = NewLogger()

type LoggerOptionFunc func(options *LoggerOption)

type LoggerOption struct {
	isDebug bool
	level string
	prefix string
}

func WithIsDebug(isDebug bool) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.isDebug = isDebug
	}
}

func WithLevel(level string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.level = level
	}
}

func WithPrefix(prefix string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.prefix = prefix
	}
}

func NewLogger(opts ...LoggerOptionFunc) go_interface_logger.InterfaceLogger {
	option := LoggerOption{
		isDebug: false,
		level: ``,
		prefix: ``,
	}
	for _, o := range opts {
		o(&option)
	}

	var logger go_interface_logger.InterfaceLogger
	if option.isDebug {
		log4go := &Log4goClass{}
		level := `debug`
		if option.level != `` {
			level = option.level
		}
		log4go.Init(option.prefix, level)
		logger = log4go
	} else {
		level := `info`
		if option.level != `` {
			level = option.level
		}
		zap := &ZapClass{}
		zap.MustInit(option.prefix, level)
		logger = zap
	}
	return logger
}

func (l *LoggerClass) Close() {
	if l.logger != nil {
		l.logger.Close()
	}
}

func (l *LoggerClass) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *LoggerClass) DebugF(format string, args ...interface{}) {
	l.logger.DebugF(format, args...)
}

func (l *LoggerClass) Print(args ...interface{}) {
	l.Info(args)
}

func (l *LoggerClass) Println(args ...interface{}) {
	l.Info(args)
}

func (l *LoggerClass) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LoggerClass) Printf(format string, args ...interface{}) {
	l.InfoF(format, args)
}
func (l *LoggerClass) InfoF(format string, args ...interface{}) {
	l.logger.InfoF(format, args...)
}

func (l *LoggerClass) Warn(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *LoggerClass) WarnF(format string, args ...interface{}) {
	l.logger.WarnF(format, args...)
}

func (l *LoggerClass) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *LoggerClass) ErrorF(format string, args ...interface{}) {
	l.logger.ErrorF(format, args...)
}

func (l *LoggerClass) ErrorWithStack(args ...interface{}) {
	l.logger.ErrorWithStack(args...)
}

func (l *LoggerClass) ErrorWithStackF(format string, args ...interface{}) {
	l.logger.ErrorWithStackF(format, args...)
}

func (l *LoggerClass) NewWriter() *Writer {
	return &Writer{}
}

type Writer struct {
}

func (l *Writer) Write(p []byte) (n int, err error) {
	Logger.Debug(string(p))
	return len(p), nil
}
