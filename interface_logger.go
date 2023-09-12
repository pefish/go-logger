package go_logger

import (
	"fmt"
	"runtime/debug"
)

type InterfaceLogger interface {
	Close()
	IsDev() bool
	IsDebug() bool
	FormatOutput(format string, args ...interface{}) string

	Debug(args ...interface{})
	DebugF(format string, args ...interface{})
	Info(args ...interface{})
	InfoF(format string, args ...interface{})
	Warn(args ...interface{})
	WarnF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
}

type loggerImpl struct {
}

var DefaultLogger = &loggerImpl{}

func (l *loggerImpl) Close() {

}

func (l *loggerImpl) IsDev() bool {
	return false
}

func (l *loggerImpl) IsDebug() bool {
	return true
}

func formatOutput(args ...interface{}) string {
	result := ``
	for _, arg := range args {
		result += fmt.Sprint(arg) + ` `
	}
	return result
}

func (l *loggerImpl) Debug(args ...interface{}) {
	fmt.Printf("[DEBUG] %s\n", formatOutput(args...))
}

func (l *loggerImpl) DebugF(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] %s\n", fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Info(args ...interface{}) {
	fmt.Printf("[INFO] %s\n", formatOutput(args...))
}

func (l *loggerImpl) InfoF(format string, args ...interface{}) {
	fmt.Printf("[INFO] %s\n", fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Warn(args ...interface{}) {
	fmt.Printf("[WARN] %s\n", formatOutput(args...))
}

func (l *loggerImpl) WarnF(format string, args ...interface{}) {
	fmt.Printf("[WARN] %s\n", fmt.Sprintf(format, args...))
}

func (l *loggerImpl) Error(args ...interface{}) {
	fmt.Printf("[ERROR] %s\n%s", formatOutput(args...), string(debug.Stack()))
}

func (l *loggerImpl) ErrorF(format string, args ...interface{}) {
	fmt.Printf("[ERROR] %s\n%s", fmt.Sprintf(format, args...), string(debug.Stack()))
}
