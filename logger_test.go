package go_logger

import (
	"errors"
)

func ExampleNewLogger() {
	Logger = NewLogger(`debug`, WithPrefix(`debug`))
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Error(errors.New(`111`))

	Logger = NewLogger(`warn`, WithPrefix(`warn`))
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Warn(123)
	Logger.Error(errors.New(`111`))

	// Output:
	//
}
