package go_logger

import (
	"errors"
	"fmt"
	"testing"
)

func TestLoggerClass_Debug(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(err)
		}
	}()

	loggerInstance := Log4goClass{}
	Logger.Init(&loggerInstance, `test`, `info`)
	Logger.Debug(1.344, `62562`)
	Logger.Error(errors.New(`111`))
	var a interface{} = `11`
	fmt.Println(a.(int))
}
