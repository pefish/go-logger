package go_logger

import (
	"errors"
)

var (
	DEFAULT_NAME  = `default`
	DEFAULT_LEVEL = `debug`
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

func (this *LoggerClass) InitWithLogger(logger InterfaceLogger, name string, level string) {
	this.logger = logger
	this.logger.Init(name, level)
}

func (this *LoggerClass) Init(name string, level string) {
	this.InitWithLogger(&Log4goClass{}, DEFAULT_NAME, DEFAULT_LEVEL)
}

func (this *LoggerClass) Close() {
	if this.logger != nil {
		this.logger.Close()
	}
}

func (this *LoggerClass) InitWithConfiguration(config Configuration) {
	if config.Logger == nil {
		panic(errors.New(`logger must be initiated`))
	}
	this.logger = config.Logger
	if config.Level == `` {
		config.Level = DEFAULT_LEVEL
	}
	this.logger.Init(config.Name, config.Level)
}

func (this *LoggerClass) Debug(args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.Debug(args...)
}

func (this *LoggerClass) DebugF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
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
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.Info(args...)
}

func (this *LoggerClass) Printf(format string, args ...interface{}) {
	this.InfoF(format, args)
}
func (this *LoggerClass) InfoF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.InfoF(format, args...)
}

func (this *LoggerClass) Warn(args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.Warn(args...)
}

func (this *LoggerClass) WarnF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.WarnF(format, args...)
}

func (this *LoggerClass) Error(args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
	}
	this.logger.Error(args...)
}

func (this *LoggerClass) ErrorF(format string, args ...interface{}) {
	if this.logger == nil {
		this.Init(DEFAULT_NAME, DEFAULT_LEVEL)
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
