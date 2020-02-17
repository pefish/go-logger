package go_logger

import (
	"fmt"
	"go.uber.org/zap"
)

type ZapClass struct {
	BaseLogger
	logger *zap.Logger
}

func (this *ZapClass) MustInit(name string, level string) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	this.logger = logger
}

func (this *ZapClass) Close() {
	this.logger.Sync()
}

func (this *ZapClass) Debug(args ...interface{}) {
	this.logger.Debug(this.FormatOutput(args...))
}

func (this *ZapClass) DebugF(format string, args ...interface{}) {
	this.logger.Debug(fmt.Sprintf(format, args...))
}

func (this *ZapClass) Info(args ...interface{}) {
	this.logger.Info(this.FormatOutput(args...))
}

func (this *ZapClass) InfoF(format string, args ...interface{}) {
	this.logger.Info(fmt.Sprintf(format, args...))
}

func (this *ZapClass) Warn(args ...interface{}) {
	this.logger.Warn(this.FormatOutput(args...))
}

func (this *ZapClass) WarnF(format string, args ...interface{}) {
	this.logger.Warn(fmt.Sprintf(format, args...))
}

func (this *ZapClass) Error(args ...interface{}) {
	this.logger.Error(this.FormatOutput(args...))
}

func (this *ZapClass) ErrorF(format string, args ...interface{}) {
	this.logger.Error(fmt.Sprintf(format, args...))
}