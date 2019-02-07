package p_logger

import (
	"gitee.com/pefish/p-go-reflect"
)

type BaseLogger struct {
}

func (this *BaseLogger) FormatOutput(args ...interface{}) string {
	result := ``
	for _, arg := range args {
		result += p_reflect.Reflect.ToString(arg) + ` `
	}
	return result
}
