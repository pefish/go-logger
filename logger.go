package go_logger

import (
	"fmt"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	t_logger "github.com/pefish/go-interface/t-logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerType struct {
	opts      *LoggerOption
	logger    *zap.Logger
	zapConfig zap.Config

	prefix string
	sync.Mutex
}

var Logger = NewLogger(t_logger.Level_INFO)

type LoggerOptionFunc func(options *LoggerOption)

type LoggerOption struct {
	printEncoding string
	level         t_logger.Level
	prefix        string
	outputFile    string // 日志输出文件
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

var errLevels = map[t_logger.Level]zapcore.Level{
	t_logger.Level_DEBUG: zap.DebugLevel,
	t_logger.Level_INFO:  zap.InfoLevel,
	t_logger.Level_WARN:  zap.WarnLevel,
	t_logger.Level_ERROR: zap.ErrorLevel,
}

func NewLogger(level t_logger.Level, opts ...LoggerOptionFunc) *LoggerType {
	option := LoggerOption{
		level:  level,
		prefix: ``,
	}
	for _, o := range opts {
		o(&option)
	}
	return newLogger(&option)
}

func newLogger(opts *LoggerOption) *LoggerType {
	isDev := false
	if opts.level != `error` && opts.level != `warn` {
		isDev = true
	}

	printEncoding := "console"
	if !isDev {
		printEncoding = "json"
	}
	outputPaths := []string{"stdout"}
	if opts.outputFile != "" {
		outputPaths = append(outputPaths, opts.outputFile)
	}
	zapConfig := zap.Config{
		DisableCaller:     true,
		DisableStacktrace: true,
		Level:             zap.NewAtomicLevelAt(errLevels[opts.level]),
		Development:       isDev,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: printEncoding,
		EncoderConfig: func() zapcore.EncoderConfig {
			if isDev {
				return zap.NewDevelopmentEncoderConfig()
			} else {
				return zap.NewProductionEncoderConfig()
			}
		}(),
		OutputPaths:      outputPaths,
		ErrorOutputPaths: []string{"stderr"},
	}
	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	return &LoggerType{
		opts:   opts,
		logger: logger,
		prefix: func(prefix string) string {
			if prefix != "" {
				return fmt.Sprintf("[%s]: ", prefix)
			} else {
				return ""
			}
		}(opts.prefix),
		zapConfig: zapConfig,
	}
}

func (l *LoggerType) CloneWithPrefix(prefix string) *LoggerType {
	defer l.Unlock()
	l.Lock()
	l.opts.prefix = prefix
	return newLogger(l.opts)
}

func (l *LoggerType) CloneWithLevel(level t_logger.Level) *LoggerType {
	defer l.Unlock()
	l.Lock()
	l.opts.level = level
	return newLogger(l.opts)
}

func (l *LoggerType) CloneWithOutputFile(filepath string) *LoggerType {
	defer l.Unlock()
	l.Lock()
	l.opts.outputFile = filepath
	return newLogger(l.opts)
}

func (l *LoggerType) Opts() *LoggerOption {
	return l.opts
}

func (l *LoggerType) Level() t_logger.Level {
	return l.opts.level
}

func (l *LoggerType) FormatOutput(args ...interface{}) string {
	return l.formatOutput("%+v", args...)
}

// 更加全面
func (l *LoggerType) Sdump(args ...interface{}) string {
	return spew.Sdump(args...)
}

func (l *LoggerType) formatOutput(format string, args ...interface{}) string {
	if len(args) == 0 {
		return ""
	}
	result := ``
	for _, arg := range args {
		result += fmt.Sprintf(format, arg) + "   "
	}
	result = result[:len(result)-3]
	return result
}

func (l *LoggerType) Debug(args ...interface{}) {
	l.logger.Debug(fmt.Sprintf("%s%s", l.prefix, l.FormatOutput(args...)))
}

func (l *LoggerType) DebugF(format string, args ...interface{}) {
	l.logger.Debug(fmt.Sprintf("%s%s", l.prefix, fmt.Sprintf(format, args...)))
}

func (l *LoggerType) DebugFRaw(format string, args ...interface{}) {
	level := l.Opts().level
	if level == "info" || level == "warn" || level == "error" {
		return
	}
	fmt.Printf("DEBUG\t%s\n", fmt.Sprintf(format, args...))
}

func (l *LoggerType) Info(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, l.FormatOutput(args...))
	l.logger.Info(msg)
}

func (l *LoggerType) InfoDump(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, l.Sdump(args...))
	l.logger.Info(msg)
}

func (l *LoggerType) InfoF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, fmt.Sprintf(format, args...))
	l.logger.Info(msg)
}

// 只支持 console 格式
func (l *LoggerType) InfoFWithRewrite(format string, args ...interface{}) {
	fmt.Printf("\r"+time.Now().Format("2006-01-02T15:04:05.000Z0700")+"\t"+zap.NewAtomicLevelAt(errLevels["info"]).Level().CapitalString()+"\t"+format, args...)
}

// 只支持 console 格式
func (l *LoggerType) InfoFRaw(format string, args ...interface{}) {
	level := l.Opts().level
	if level == "warn" || level == "error" {
		return
	}
	fmt.Printf("INFO\t%s\n", fmt.Sprintf(format, args...))
}

func (l *LoggerType) Warn(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, l.FormatOutput(args...))
	l.logger.Warn(msg)
}

func (l *LoggerType) WarnF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, fmt.Sprintf(format, args...))
	l.logger.Warn(msg)
}

func (l *LoggerType) WarnFRaw(format string, args ...interface{}) {
	level := l.Opts().level
	if level == "error" {
		return
	}
	fmt.Printf("WARN\t%s\n", fmt.Sprintf(format, args...))
}

func (l *LoggerType) Error(args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, l.FormatOutput(args...))
	l.logger.Error(msg)
}

func (l *LoggerType) ErrorF(format string, args ...interface{}) {
	msg := fmt.Sprintf("%s%s", l.prefix, fmt.Sprintf(format, args...))
	l.logger.Error(msg)
}

func (l *LoggerType) ErrorFRaw(format string, args ...interface{}) {
	fmt.Printf("ERROR\t%s\n", fmt.Sprintf(format, args...))
}
