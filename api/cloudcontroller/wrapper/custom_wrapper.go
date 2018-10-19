package wrapper

import (
	"code.cloudfoundry.org/cli/api/shared"
)

// CustomWrapper is a wrapper that can execute arbitrary code via the
// CustomMake function on every request that passes through Make.
type CustomWrapper struct {
	connection shared.Connection
	CustomMake func(connection shared.Connection, request *shared.Request, passedResponse shared.Response) error
}

func (e *CustomWrapper) Make(request *shared.Request, passedResponse shared.Response) error {
	return e.CustomMake(e.connection, request, passedResponse)
}

func (e *CustomWrapper) Wrap(innerconnection shared.Connection) shared.Connection {
	e.connection = innerconnection
	return e
}
