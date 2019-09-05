package go_logger

import (
	"errors"
)

type LoggerClass struct {
	logger InterfaceLogger
}

type Configuration struct {
	Logger InterfaceLogger
	Name   string
	Level  string
}

var Logger = &LoggerClass{}

func (this *LoggerClass) Init(logger InterfaceLogger, name string, level string) {
	this.logger = logger
	this.logger.Init(name, level)
}

func (this *LoggerClass) InitWithConfiguration(config Configuration) {
	if config.Logger == nil {
		panic(errors.New(`logger must be initiated`))
	}
	this.logger = config.Logger
	if config.Level == `` {
		config.Level = `info`
	}
	this.logger.Init(config.Name, config.Level)
}

func (this *LoggerClass) Debug(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.Debug(args...)
}

func (this *LoggerClass) DebugF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.DebugF(format, args...)
}

func (this *LoggerClass) Print(args ...interface{}) {
	this.Info(args)
}

func (this *LoggerClass) Println(args ...interface{}) {
	this.Info(args)
}

func (this *LoggerClass) Info(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.Info(args...)
}

func (this *LoggerClass) Printf(format string, args ...interface{}) {
	this.InfoF(format, args)
}
func (this *LoggerClass) InfoF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.InfoF(format, args...)
}

func (this *LoggerClass) Warn(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.Warn(args...)
}

func (this *LoggerClass) WarnF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.WarnF(format, args...)
}

func (this *LoggerClass) Error(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
	this.logger.Error(args...)
}

func (this *LoggerClass) ErrorF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`, `info`)
	}
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


