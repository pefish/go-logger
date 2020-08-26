package go_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapClass struct {
	logger *zap.Logger
	prefix string
}

var Logger = NewLogger("info")

type LoggerOptionFunc func(options *LoggerOption)

type LoggerOption struct {
	printEncoding string
	level       string
	prefix      string
	isDev bool
}

func WithIsDev(isDev bool) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.isDev = isDev
	}
}

func WithPrintEncoding(printEncoding string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.printEncoding = printEncoding
	}
}

func WithPrefix(prefix string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.prefix = prefix
	}
}

var errLevels = map[string]zapcore.Level{
	`debug`: zap.DebugLevel,
	`info`:  zap.InfoLevel,
	`warn`:  zap.WarnLevel,
	`error`: zap.ErrorLevel,
}

func NewLogger(level string, opts ...LoggerOptionFunc) *ZapClass {
	option := LoggerOption{
		level:  level,
		prefix: ``,
		isDev: false,
	}

	if option.level != `error` && option.level != `warn` {
		option.isDev = true
	}

	for _, o := range opts {
		o(&option)
	}

	if !option.isDev {
		option.printEncoding = "json"
	} else {
		option.printEncoding = "console"
	}
	//fmt.Printf("%#v\n", option)

	logger, err := zap.Config{
		DisableCaller: true,
		DisableStacktrace: true,
		Level:       zap.NewAtomicLevelAt(errLevels[option.level]),
		Development: option.isDev,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         option.printEncoding,
		EncoderConfig: func() zapcore.EncoderConfig {
			if option.isDev {
				return zap.NewDevelopmentEncoderConfig()
			} else {
				return zap.NewProductionEncoderConfig()
			}
		}(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		panic(err)
	}
	return &ZapClass{
		logger: logger,
		prefix: func() string {
			if option.prefix != "" {
				return fmt.Sprintf("[%s]: ", option.prefix)
			} else {
				return ""
			}
		}(),
	}
}

func (zapInstance *ZapClass) Close() {
	zapInstance.logger.Sync()
}

func (zapInstance *ZapClass) FormatOutput(format string, args ...interface{}) string {
	result := ``
	for _, arg := range args {
		result += fmt.Sprintf(format, arg) + "   "
	}
	result = result[:len(result) - 3]
	return result
}

func (zapInstance *ZapClass) Debug(args ...interface{}) {
	zapInstance.logger.Debug(fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput("%v", args...)))
}

func (zapInstance *ZapClass) DebugF(format string, args ...interface{}) {
	zapInstance.logger.Debug(fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...)))
}

func (zapInstance *ZapClass) Info(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput("%v", args...))
	zapInstance.logger.Info(msg)
}

func (zapInstance *ZapClass) InfoF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Info(msg)
}

func (zapInstance *ZapClass) Warn(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput("%v", args...))
	zapInstance.logger.Warn(msg)
}

func (zapInstance *ZapClass) WarnF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Warn(msg)
}

func (zapInstance *ZapClass) Error(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, zapInstance.FormatOutput("%+v", args...))
	zapInstance.logger.Error(msg)
}

func (zapInstance *ZapClass) ErrorF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", zapInstance.prefix, fmt.Sprintf(format, args...))
	zapInstance.logger.Error(msg)
}
