package go_logger

import (
	"github.com/pefish/go-application"
	"github.com/pefish/go-interface-logger"
)

type LoggerClass struct {
	logger go_interface_logger.InterfaceLogger
}


var Logger = NewLogger()

type LoggerOptionFunc func(options *LoggerOption)

type LoggerOption struct {
	name string
	level string
}

func WithName(name string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.name = name
	}
}

func WithLevel(level string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.level = level
	}
}

func NewLogger(opts ...LoggerOptionFunc) go_interface_logger.InterfaceLogger {
	option := LoggerOption{
		name: `default`,
		level: `debug`,
	}
	for _, o := range opts {
		o(&option)
	}

	var logger go_interface_logger.InterfaceLogger
	if go_application.Application.Debug {
		log4go := &Log4goClass{}
		log4go.Init(option.name, option.level)
		logger = log4go
	} else {
		//logrus := &LogrusClass{}
		//logrus.Init(option.name, option.level)
		//logger = logrus
		zap := &ZapClass{}
		zap.MustInit(option.name, option.level)
		logger = zap
	}
	return logger
}

func (this *LoggerClass) Close() {
	if this.logger != nil {
		this.logger.Close()
	}
}

func (this *LoggerClass) Debug(args ...interface{}) {
	this.logger.Debug(args...)
}

func (this *LoggerClass) DebugF(format string, args ...interface{}) {
	this.logger.DebugF(format, args...)
}

func (this *LoggerClass) Print(args ...interface{}) {
	this.Info(args)
}

func (this *LoggerClass) Println(args ...interface{}) {
	this.Info(args)
}

func (this *LoggerClass) Info(args ...interface{}) {
	this.logger.Info(args...)
}

func (this *LoggerClass) Printf(format string, args ...interface{}) {
	this.InfoF(format, args)
}
func (this *LoggerClass) InfoF(format string, args ...interface{}) {
	this.logger.InfoF(format, args...)
}

func (this *LoggerClass) Warn(args ...interface{}) {
	this.logger.Warn(args...)
}

func (this *LoggerClass) WarnF(format string, args ...interface{}) {
	this.logger.WarnF(format, args...)
}

func (this *LoggerClass) Error(args ...interface{}) {
	this.logger.Error(args...)
}

func (this *LoggerClass) ErrorF(format string, args ...interface{}) {
	this.logger.ErrorF(format, args...)
}

func (this *LoggerClass) NewWriter() *Writer {
	return &Writer{}
}

type Writer struct {
}

func (this *Writer) Write(p []byte) (n int, err error) {
	Logger.Debug(string(p))
	return len(p), nil
}
