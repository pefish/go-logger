package go_logger

import (
	"errors"
	"fmt"
)

func ExampleNewLogger() {
	defer func() {
		if err := recover(); err != nil {
			Logger.Error(err)
		}
	}()

	Logger = NewLogger(WithName(`haha`))
	Logger.Debug(1.344, `62562`)
	Logger.InfoF(`hs%sfdga%s`, `6245w`, `111`)
	Logger.Error(errors.New(`111`))
	var a interface{} = `11`
	fmt.Println(a.(int))

	// Output:
	// {"level":"info","ts":1581927455.8056371,"caller":"go-logger/zap.go:38","msg":"hs6245wfdga111"}
	// {"level":"error","ts":1581927455.805695,"caller":"go-logger/zap.go:50","msg":"111 ","stacktrace":"github.com/pefish/go-logger.(*ZapClass).Error\n\t/Users/joy/Work/backend/go-logger/zap.go:50\ngithub.com/pefish/go-logger.ExampleNewLogger\n\t/Users/joy/Work/backend/go-logger/logger_test.go:20\ntesting.runExample\n\t/usr/local/go/src/testing/run_example.go:62\ntesting.runExamples\n\t/usr/local/go/src/testing/example.go:44\ntesting.(*M).Run\n\t/usr/local/go/src/testing/testing.go:1118\nmain.main\n\t_testmain.go:44\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
	// {"level":"error","ts":1581927455.805763,"caller":"go-logger/zap.go:50","msg":"interface conversion: interface {} is string, not int ","stacktrace":"github.com/pefish/go-logger.(*ZapClass).Error\n\t/Users/joy/Work/backend/go-logger/zap.go:50\ngithub.com/pefish/go-logger.ExampleNewLogger.func1\n\t/Users/joy/Work/backend/go-logger/logger_test.go:12\nruntime.gopanic\n\t/usr/local/go/src/runtime/panic.go:679\nruntime.panicdottypeE\n\t/usr/local/go/src/runtime/iface.go:255\ngithub.com/pefish/go-logger.ExampleNewLogger\n\t/Users/joy/Work/backend/go-logger/logger_test.go:22\ntesting.runExample\n\t/usr/local/go/src/testing/run_example.go:62\ntesting.runExamples\n\t/usr/local/go/src/testing/example.go:44\ntesting.(*M).Run\n\t/usr/local/go/src/testing/testing.go:1118\nmain.main\n\t_testmain.go:44\nruntime.main\n\t/usr/local/go/src/runtime/proc.go:203"}
}
