package go_logger

import (
	"errors"
	"fmt"
	"github.com/pefish/go-application"
	"testing"
)

func TestLoggerClass_Debug(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(err)
		}
	}()

	go_application.Application.Debug = false
	Logger.Init(`haha`, `debug`)
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Error(errors.New(`111`))
	var a interface{} = `11`
	fmt.Println(a.(int))
}
