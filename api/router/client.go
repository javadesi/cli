// Package router is a GoLang library that interacts with CloudFoundry Go Router
package router

import (
	"fmt"
	"runtime"

	"github.com/tedsuo/rata"
)

// Client is a client that can be used to talk to a Cloud Controller's V2
// Endpoints.
type Client struct {
	routerGroupEndpoint string

	connection Connection
	router     *rata.RequestGenerator
	userAgent  string
	wrappers   []ConnectionWrapper
}

// ClientConfig allows the Client to be configured
type ClientConfig struct {
	// AppName is the name of the application/process using the client.
	AppName string

	// AppVersion is the version of the application/process using the client.
	AppVersion string

	// Wrappers that apply to the client connection.
	Wrappers []ConnectionWrapper
}

// NewClient returns a new Router Client.
func NewClient(config ClientConfig) *Client {
	userAgent := fmt.Sprintf("%s/%s (%s; %s %s)",
		config.AppName,
		config.AppVersion,
		runtime.Version(),
		runtime.GOARCH,
		runtime.GOOS,
	)

	client := Client{
		userAgent: userAgent,
	}

	// TODO Wrap Connection
	// client.WrapConnection(NewErrorWrapper())

	return &client
}
