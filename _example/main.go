package main

import (
	go_logger "github.com/pefish/go-logger"
	"github.com/pkg/errors"
	"time"
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
	go_logger.Logger = go_logger.NewLogger("debug")
	go_logger.Logger.Info("haha")
	for i := 0; i < 100; i++ {
		go_logger.Logger.InfoFWithRewrite("%d", i)
		//fmt.Printf("\r111%d", i)
		time.Sleep(time.Second)
	}
}

