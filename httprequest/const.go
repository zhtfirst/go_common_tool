package httprequest

import (
	"errors"
	"net/http"
)

const (
	pathKey   = "path"
	formKey   = "form"
	headerKey = "header"
	jsonKey   = "json"
	slash     = "/"
	colon     = ':'
	trace     = "traceid"
	// ContentType is the header key for Content-Type.
	ContentType = "Content-Type"
	// JsonContentType is the content type for JSON.
	JsonContentType = "application/json; charset=utf-8"
	FormContentType = "application/x-www-form-urlencoded"
	responseLength  = 2048 //日志response的长度

	// defaultTimeout is the default timeout in seconds.
	defaultTimeout = 30
	// defaultServiceName is the default service name.
	defaultServiceName = "http_request"
)

var (
	supportMethod = map[string]bool{
		http.MethodGet:     true,
		http.MethodPost:    true,
		http.MethodConnect: true,
		http.MethodDelete:  true,
		http.MethodHead:    true,
		http.MethodPatch:   true,
		http.MethodPut:     true,
		http.MethodTrace:   true,
	}
	noSupportMethod = errors.New("current request method is not supported")
)
