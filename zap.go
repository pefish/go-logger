package go_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapClass struct {
	BaseLogger
	logger *zap.Logger
}

var errLevels = map[string]zapcore.Level{
	`debug`: zap.DebugLevel,
	`info`: zap.InfoLevel,
	`warn`: zap.WarnLevel,
	`error`: zap.ErrorLevel,
}

func (this *ZapClass) MustInit(name string, level string) {
	// 生产环境必须info级别或以上
	if level != `error` && level != `warn` && level != `info` {
		panic(`level error`)
	}
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(errLevels[level]),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    zap.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
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
	msg := this.FormatOutput(args...)
	this.logger.Info(msg, zap.String("message", msg), zap.String("severity", "info"))
}

func (this *ZapClass) InfoF(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.logger.Info(msg, zap.String("message", msg), zap.String("severity", "info"))
}

func (this *ZapClass) Warn(args ...interface{}) {
	msg := this.FormatOutput(args...)
	this.logger.Warn(msg, zap.String("message", msg), zap.String("severity", "warning"))
}

func (this *ZapClass) WarnF(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.logger.Warn(msg, zap.String("message", msg), zap.String("severity", "warning"))
}

func (this *ZapClass) Error(args ...interface{}) {
	msg := this.FormatOutput(args...)
	this.logger.Error(msg, zap.String("message", msg), zap.String("severity", "error"))
}

func (this *ZapClass) ErrorF(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	this.logger.Error(msg, zap.String("message", msg), zap.String("severity", "error"))
}