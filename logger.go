package go_logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapClass struct {
	opts *LoggerOption
	logger *zap.Logger
	isDev bool  // 日志级别不是error、warn，则为开发模式
	isDebug bool  // 日志级别不是error、warn、info，则为开发模式

	prefix string
}

var Logger = NewLogger("info")

type LoggerOptionFunc func(options *LoggerOption)

type LoggerOption struct {
	printEncoding string
	level       string
	prefix      string
	outputFile string  // 日志输出文件
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

func WithOutputFile(filepath string) LoggerOptionFunc {
	return func(option *LoggerOption) {
		option.outputFile = filepath
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
	}
	for _, o := range opts {
		o(&option)
	}
	return newLogger(&option)
}

func newLogger(opts *LoggerOption) *ZapClass {
	isDev := false
	isDebug := false
	if opts.level != `error` && opts.level != `warn` {
		isDev = true
		if opts.level != `info` {
			isDebug = true
		}
	}

	printEncoding := "console"
	if !isDev {
		printEncoding = "json"
	}
	outputPaths := []string{"stdout"}
	if opts.outputFile != "" {
		outputPaths = append(outputPaths, opts.outputFile)
	}
	logger, err := zap.Config{
		DisableCaller: true,
		DisableStacktrace: true,
		Level:       zap.NewAtomicLevelAt(errLevels[opts.level]),
		Development: isDev,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         printEncoding,
		EncoderConfig: func() zapcore.EncoderConfig {
			if isDev {
				return zap.NewDevelopmentEncoderConfig()
			} else {
				return zap.NewProductionEncoderConfig()
			}
		}(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		panic(err)
	}
	return &ZapClass{
		opts: opts,
		logger: logger,
		prefix: func() string {
			if opts.prefix != "" {
				return fmt.Sprintf("[%s]: ", opts.prefix)
			} else {
				return ""
			}
		}(),
		isDev: isDev,
		isDebug: isDebug,
	}
}

func (zapInstance *ZapClass) CloneWithPrefix(prefix string) *ZapClass {
	zapInstance.opts.prefix = prefix
	return newLogger(zapInstance.opts)
}

func (zapInstance *ZapClass) CloneWithLevel(level string) *ZapClass {
	zapInstance.opts.level = level
	return newLogger(zapInstance.opts)
}

func (zapInstance *ZapClass) Close() {
	zapInstance.logger.Sync()
}

func (zapInstance *ZapClass) IsDev() bool {
	return zapInstance.isDev
}

func (zapInstance *ZapClass) Opts() *LoggerOption {
	return zapInstance.opts
}

func (zapInstance *ZapClass) IsDebug() bool {
	return zapInstance.isDebug
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
