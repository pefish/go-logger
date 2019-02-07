package p_logger

import (
	"testing"
)

func TestLoggerClass_Debug(t *testing.T) {
	loggerInstance := Log4goClass{}
	Logger.Init(&loggerInstance, `test`)
	Logger.Debug(1.344, `62562`)
}
