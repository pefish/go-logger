package go_logger

type InterfaceLogger interface {
	Close()
	IsDev() bool
	IsDebug() bool
	FormatOutput(args ...interface{}) string

	Debug(args ...interface{})
	DebugF(format string, args ...interface{})
	Info(args ...interface{})
	InfoF(format string, args ...interface{})
	Warn(args ...interface{})
	WarnF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
}
