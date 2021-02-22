package go_logger

import (
	"fmt"
	"github.com/pkg/errors"
)

func ExampleNewLogger() {
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(err)
		}
	}()

	Logger = NewLogger(`debug`, WithPrefix(`debug`))
	fmt.Println(Logger.IsDev())
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Error(errors.New(`111`))

	Logger = NewLogger(`warn`, WithPrefix(`warn`))
	fmt.Println(Logger.IsDev())
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Warn(123)
	Logger.Error(errors.New(`111`))

	logger1 := Logger.CloneWithPrefix("haha").CloneWithLevel("debug")
	logger1.Debug(1.344, `debug`)
	logger1.Warn(1.344, `warn`)
	// Output:
	// true
	// false
}

