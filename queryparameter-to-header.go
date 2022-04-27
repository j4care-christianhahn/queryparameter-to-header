// Package plugindemo a demo plugin.
package queryparameter_to_header

import (
	"context"
	"fmt"
	"net/http"
)

// the plugin configuration.
type Config struct {
	QueryParameter string `json:"query_parameter"`
	Header         string `json:"header"`
}

// CreateConfig creates the default plugin configuration
func CreateConfig() *Config {
	return &Config{
		QueryParameter: "v",
		Header:         "X-Version",
	}
}

type QueryParameterToHeaderMiddleware struct {
	next           http.Handler
	queryParameter string
	header         string
	name           string
}

// Creates a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if len(config.Header) < 1 {
		return nil, fmt.Errorf("header cannot be empty string")
	}
	if len(config.QueryParameter) < 1 {
		return nil, fmt.Errorf("query parameter cannot be empty string")
	}

	return &QueryParameterToHeaderMiddleware{
		header:         config.Header,
		queryParameter: config.QueryParameter,
		next:           next,
		name:           name,
	}, nil
}

func (m *QueryParameterToHeaderMiddleware) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	parameterValues := query[m.queryParameter]
	if len(parameterValues) > 0 {
		req.Header.Set(m.header, parameterValues[0])
	}
	m.next.ServeHTTP(rw, req)
}
