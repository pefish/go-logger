package p_logger

import (
	"github.com/pefish/go-application"
)

type LoggerClass struct {
	logger InterfaceLogger
}

var Logger = &LoggerClass{}

func (this *LoggerClass) Init(logger InterfaceLogger, name string) {
	this.logger = logger
	this.logger.Init(name, p_application.Application.Debug)
}

func (this *LoggerClass) Debug(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`)
	}
	this.logger.Debug(args...)
}

func (this *LoggerClass) Info(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`)
	}
	this.logger.Info(args...)
}

func (this *LoggerClass) Warn(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`)
	}
	this.logger.Warn(args...)
}

func (this *LoggerClass) Error(args ...interface{}) {
	if this.logger == nil {
		this.Init(&Log4goClass{}, `default`)
	}
	this.logger.Error(args...)
}