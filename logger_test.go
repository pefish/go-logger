package go_logger

import (
	"testing"

	"github.com/pkg/errors"
)

func TestLogger(t *testing.T) {
	logger := NewLogger(`debug`, WithPrefix(`debug`))
	logger.Debug(1.344, `62562`)
	logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	logger.Error(errors.New(`111`))

	logger = NewLogger(`warn`, WithPrefix(`warn`))
	logger.Debug(1.344, `62562`)
	logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	logger.Warn(123)
	logger.Error(errors.New(`111`))

	logger1 := logger.CloneWithPrefix("haha").CloneWithLevel("debug")
	logger1.Debug(1.344, `debug`)
	logger1.Warn(1.344, `warn`)

	logger2 := logger.CloneWithPrefix("xixi").CloneWithLevel("debug")
	logger2.Debug(1.344, `debug`)
	logger2.Warn(1.344, `warn`)

	logger.InfoFRaw("gsfga")
	logger.ErrorFRaw("test")
}
