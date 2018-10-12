// Package router contains utilities to make call to the router API
package router

import "net/http"

//go:generate counterfeiter . Connection

// Connection creates and executes http requests
type Connection interface {
	Make(request *http.Request, passedResponse *Response) error
}
