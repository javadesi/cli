package router

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
)

// ConnectionConfig is for configuring the RouterConnection
type ConnectionConfig struct {
	DialTimeout       time.Duration
	SkipSSLValidation bool
}

// RouterConnection represents the connection to Router
type RouterConnection struct {
	HTTPClient *http.Client
}

// NewConnection returns a pointer to a new UAA Connection with the provided configuration
func NewConnection(config ConnectionConfig) *RouterConnection {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.SkipSSLValidation,
		},
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   config.DialTimeout,
		}).DialContext,
	}

	return &RouterConnection{
		HTTPClient: &http.Client{Transport: tr},
	}
}

// Make performs the request and parses the response.
func (connection *RouterConnection) Make(request *Request, passedResponse *Response) error {
	// In case this function is called from a retry, passedResponse may already
	// be populated with a previous response. We reset in case there's an HTTP
	// error and we don't repopulate it in populateResponse.
	passedResponse.reset()

	response, err := connection.HTTPClient.Do(request.Request)
	if err != nil {
		return connection.processRequestErrors(request.Request, err)
	}

	return connection.populateResponse(response, passedResponse)
}

func (*RouterConnection) handleStatusCodes(response *http.Response, passedResponse *Response) error {
	if response.StatusCode == http.StatusNoContent {
		passedResponse.RawResponse = []byte("{}")
	} else {
		rawBytes, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return err
		}

		passedResponse.RawResponse = rawBytes
	}

	if response.StatusCode >= 400 {
		return ccerror.RawHTTPStatusError{
			StatusCode:  response.StatusCode,
			RawResponse: passedResponse.RawResponse,
			RequestIDs:  response.Header["X-Vcap-Request-Id"],
		}
	}

	return nil
}
func (connection *RouterConnection) populateResponse(response *http.Response, passedResponse *Response) error {
	passedResponse.HTTPResponse = response

	rawBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	passedResponse.RawResponse = rawBytes

	err = connection.handleStatusCodes(response, passedResponse)
	if err != nil {
		return errConfig
	}

	if passedResponse.Result != nil {
		decoder := json.NewDecoder(bytes.NewBuffer(passedResponse.RawResponse))
		decoder.UseNumber()
		err = decoder.Decode(passedResponse.Result)
		if err != nil {
			return err
		}
	}

	return nil
}

func (*RouterConnection) processRequestErrors(request *http.Request, err error) error {
	switch e := err.(type) {
	case *url.Error:
		switch urlErr := e.Err.(type) {
		case x509.UnknownAuthorityError:
			return ccerror.UnverifiedServerError{
				URL: request.URL.String(),
			}
		case x509.HostnameError:
			return ccerror.SSLValidationHostnameError{
				Message: urlErr.Error(),
			}
		default:
			return ccerror.RequestError{Err: e}
		}
	default:
		return err
	}
}
