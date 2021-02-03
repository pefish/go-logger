package main

import (
	go_logger "github.com/pefish/go-logger"
	"github.com/pkg/errors"
)

func test() error {
	return test3()
}

func test3() error {
	return test1()
}

func test1() error {
	return errors.New("13451")
}

func main() {
	//err := test()
	//if err != nil {
	//	go_logger.Logger.Error(err)
	//}
	go_logger.Logger = go_logger.NewLogger("debug", go_logger.WithOutputFile("/Users/yunchuang/Work/backend/golang/go-logger/1.log"))
	go_logger.Logger.Info(111)
	go_logger.Logger.Error(errors.New("123"))
}