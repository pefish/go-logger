package go_logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogrusClass struct {
	BaseLogger
	logger *logrus.Entry
}

func (l *LogrusClass) Init(name string, level string) {
	logrus.SetFormatter(&FluentdFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	logrus.AddHook(Hook{
		mu:       &sync.Mutex{},
		file:     true,
		line:     true,
		function: true,
		levels:   logrus.AllLevels,
	})
	l.logger = logrus.WithFields(logrus.Fields{
		"project": name,
	})
}

func (l *LogrusClass) Close() {

}

func (l *LogrusClass) Debug(args ...interface{}) {
	l.logger.Debugln(l.FormatOutput(args...))
}

func (l *LogrusClass) DebugF(format string, args ...interface{}) {
	l.logger.Debugln(fmt.Sprintf(format, args...))
}

func (l *LogrusClass) Info(args ...interface{}) {
	l.logger.Infoln(l.FormatOutput(args...))
}

func (l *LogrusClass) InfoF(format string, args ...interface{}) {
	l.logger.Infoln(fmt.Sprintf(format, args...))
}

func (l *LogrusClass) Warn(args ...interface{}) {
	l.logger.Warnln(l.FormatOutput(args...))
}

func (l *LogrusClass) WarnF(format string, args ...interface{}) {
	l.logger.Warnln(fmt.Sprintf(format, args...))
}

func (l *LogrusClass) Error(args ...interface{}) {
	l.logger.Errorln(l.FormatOutput(args...))
}

func (l *LogrusClass) ErrorF(format string, args ...interface{}) {
	l.logger.Errorln(fmt.Sprintf(format, args...))
}

// -------------------------------- FluentdFormatter --------------------------------

type FluentdFormatter struct {
	TimestampFormat string
}

// Format the log entry. Implements logrus.Formatter.
func (f *FluentdFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			// https://github.com/Sirupsen/logrus/issues/137
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}
	prefixFieldClashes(data)

	timestampFormat := f.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = time.RFC3339Nano
	}

	data["time"] = entry.Time.Format(timestampFormat)
	data["message"] = entry.Message
	data["severity"] = entry.Level.String()

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

func prefixFieldClashes(data logrus.Fields) {
	if t, ok := data["time"]; ok {
		data["fields.time"] = t
	}

	if m, ok := data["msg"]; ok {
		data["fields.msg"] = m
	}

	if l, ok := data["level"]; ok {
		data["fields.level"] = l
	}
}

// -------------------------------- Hook --------------------------------

type Hook struct {
	mu       *sync.Mutex
	file     bool
	line     bool
	function bool
	levels   []logrus.Level
}

// Levels is ...
func (h Hook) Levels() []logrus.Level {
	return h.levels
}

// Fire is ...
func (h Hook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, 64)
	cnt := runtime.Callers(3, pc)
	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i])
		name := fu.Name()
		if !strings.Contains(name, "github.com/sirupsen/logrus") && !strings.Contains(name, `core`) {
			file, line := fu.FileLine(pc[i] - 1)
			if h.file {
				h.mu.Lock()
				entry.Data["file"] = path.Base(file)
				h.mu.Unlock()
			}

			if h.function {
				h.mu.Lock()
				entry.Data["func"] = path.Base(name)
				h.mu.Unlock()
			}

			if h.line {
				h.mu.Lock()
				entry.Data["line"] = line
				h.mu.Unlock()
			}

			break
		}
	}

	return nil
}
