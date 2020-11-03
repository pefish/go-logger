package go_logger

import (
	"fmt"
	"github.com/pkg/errors"
)

func ExampleNewLogger() {
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
	// Output:
	// true
	// false
}

