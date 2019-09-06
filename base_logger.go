package go_logger

import (
	"fmt"
)

type BaseLogger struct {
}

func (this *BaseLogger) FormatOutput(args ...interface{}) string {
	result := ``
	for _, arg := range args {
		result += fmt.Sprint(arg) + ` `
	}
	return result
}
