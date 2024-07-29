package main

import (
	i_logger "github.com/pefish/go-interface/i-logger"
	go_logger "github.com/pefish/go-logger"
	"github.com/pkg/errors"
)

func test3() error {
	return test1()
}

func test1() error {
	return errors.New("13451")
}

func main() {
	var logger1 i_logger.ILogger = go_logger.Logger.CloneWithPrefix("haha")
	logger1.Debug(1.344, `debug`)
	logger1.Warn(1.344, `warn`)

	var logger2 i_logger.ILogger = go_logger.Logger.CloneWithPrefix("xixi")
	logger2.Debug(1.344, `debug`)
	logger2.Warn(1.344, `warn`)

	logger2.Warn(
		logger2.Sdump(map[string]interface{}{
			"abc":   "abc",
			"gsdfg": 232,
		}),
	)
	logger2.InfoDump(
		map[string]interface{}{
			"abc":   "abc",
			"gsdfg": 232,
		},
	)
}
