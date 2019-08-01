package go_logger

type InterfaceLogger interface {
	Init(name string, debug bool)
	Close()

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

