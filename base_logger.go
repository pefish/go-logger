package go_logger

import (
	"github.com/pefish/go-reflect"
)

type BaseLogger struct {
}

func (this *BaseLogger) FormatOutput(args ...interface{}) string {
	result := ``
	for _, arg := range args {
		result += go_reflect.Reflect.ToString(arg) + ` `
	}
	return result
}
