package go_logger

type InterfaceLogger interface {
	Init(name string, level string)
	Close()

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

