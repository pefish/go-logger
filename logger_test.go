package go_logger

import (
	"testing"

	"github.com/pkg/errors"
)

func TestLogger(t *testing.T) {
	Logger = NewLogger(`debug`, WithPrefix(`debug`))
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Error(errors.New(`111`))

	Logger = NewLogger(`warn`, WithPrefix(`warn`))
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Warn(123)
	Logger.Error(errors.New(`111`))

	logger1 := Logger.CloneWithPrefix("haha").CloneWithLevel("debug")
	logger1.Debug(1.344, `debug`)
	logger1.Warn(1.344, `warn`)

	logger2 := Logger.CloneWithPrefix("xixi").CloneWithLevel("debug")
	logger2.Debug(1.344, `debug`)
	logger2.Warn(1.344, `warn`)

	Logger.InfoFRaw("gsfga")
	Logger.ErrorFRaw("test")
}
