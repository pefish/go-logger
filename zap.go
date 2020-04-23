package go_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapClass struct {
	BaseLogger
	logger *zap.Logger
	prefix string
}

var errLevels = map[string]zapcore.Level{
	`debug`: zap.DebugLevel,
	`info`: zap.InfoLevel,
	`warn`: zap.WarnLevel,
	`error`: zap.ErrorLevel,
}

func (zapInstance *ZapClass) MustInit(prefix string, level string) {
	if prefix != `` {
		zapInstance.prefix = fmt.Sprintf("[%s]: ", prefix)
	}
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
	zapInstance.logger = logger
}

func (zapInstance *ZapClass) Close() {
	zapInstance.logger.Sync()
}

func (zapInstance *ZapClass) Debug(args ...interface{}) {
	zapInstance.logger.Debug(fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput(args...)))
}

func (zapInstance *ZapClass) DebugF(format string, args ...interface{}) {
	zapInstance.logger.Debug(fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...)))
}

func (zapInstance *ZapClass) Info(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput(args...))
	zapInstance.logger.Info(msg, zap.String("message", msg), zap.String("severity", "info"))
}

func (zapInstance *ZapClass) InfoF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Info(msg, zap.String("message", msg), zap.String("severity", "info"))
}

func (zapInstance *ZapClass) Warn(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput(args...))
	zapInstance.logger.Warn(msg, zap.String("message", msg), zap.String("severity", "warning"))
}

func (zapInstance *ZapClass) WarnF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Warn(msg, zap.String("message", msg), zap.String("severity", "warning"))
}

func (zapInstance *ZapClass) Error(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput(args...))
	zapInstance.logger.Error(msg, zap.String("message", msg), zap.String("severity", "error"))
}

func (zapInstance *ZapClass) ErrorF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Error(msg, zap.String("message", msg), zap.String("severity", "error"))
}

func (zapInstance *ZapClass) ErrorWithStack(args ...interface{}) {
	zapInstance.Error(args...)
}

func (zapInstance *ZapClass) ErrorWithStackF(format string, args ...interface{}) {
	zapInstance.ErrorF(format, args...)
}