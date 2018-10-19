package shared

import "net/http"

type Response interface {
	PopulateFrom(*http.Response) error
	Reset()

	StatusCode() int
	// TODO: look at whether this could be renamed to HTTPResponse
	GetHTTPResponse() *http.Response
}
