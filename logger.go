package go_logger

import (
	"github.com/pefish/go-application"
)

var (
	DEFAULT_NAME  = `default`
)

type LoggerClass struct {
	logger InterfaceLogger
}

type Configuration struct {
	Name   string
	Level  string
}

var Logger = &LoggerClass{}

func (this *LoggerClass) Init(name string, level string) {
	if name == `` {
		name = DEFAULT_NAME
	}
	if level == `` {
		level = `debug`
	}
	if go_application.Application.Debug {
		this.logger = &Log4goClass{}
	} else {
		this.logger = &LogrusClass{}
		level = `info`
	}
	this.logger.Init(name, level)
}

func (this *LoggerClass) Close() {
	if this.logger != nil {
		this.logger.Close()
	}
}

func (this *LoggerClass) InitWithConfiguration(config Configuration) {
	this.logger.Init(config.Name, config.Level)
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
