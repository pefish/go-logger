# go-logger

#### Description
logger tool

#### Usage

```golang
import "github.com/pefish/go-logger"

go_logger.Logger = go_logger.NewLogger(go_logger.WithIsDebug(env != `prod`))

go_logger.Logger.DebugF(`%#v`, map[string]string{
		`haha`: `11`,
	})
go_logger.Logger.Info(`111`)
go_logger.Logger.WarnF(`%s`, `5636`)
go_logger.Logger.Error(errors.New(`error`))

```
