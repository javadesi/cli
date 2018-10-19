package shared

import "net/http"

type Response interface {
	PopulateFrom(*http.Response) error
	Reset()
}
