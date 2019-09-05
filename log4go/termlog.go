// Copyright (C) 2010, Kyle Lemons <kyle@kylelemons.net>.  All rights reserved.

package log4go

import (
	"fmt"
)

// This is the standard writer that prints to standard output.
type ConsoleLogWriter struct {
	format string
}

// This creates a new ConsoleLogWriter
func NewConsoleLogWriter() *ConsoleLogWriter {
	consoleWriter := &ConsoleLogWriter{
		format: "[%T %D] [%C] [%L] (%S) %M",
	}
	return consoleWriter
}
func (c *ConsoleLogWriter) SetFormat(format string) {
	c.format = format
}

// This is the ConsoleLogWriter's output method.  This will block if the output
// buffer is full.
func (c *ConsoleLogWriter) LogWrite(rec *LogRecord) {
	fmt.Println(FormatLogRecord(c.format, rec))
}

// Close stops the logger from sending messages to standard output.  Attempts to
// send log messages to this logger after a Close have undefined behavior.
func (c *ConsoleLogWriter) Close() {

}
