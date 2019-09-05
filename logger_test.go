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
	Logger.InitWithLogger(&loggerInstance, `haha`, `debug`)
	Logger.Debug(1.344, `62562`)
	Logger.Error(errors.New(`111`))
	Logger.DebugF(`hs%sfdga%s`, `6245w`, `111`)
	var a interface{} = `11`
	fmt.Println(a.(int))
}
