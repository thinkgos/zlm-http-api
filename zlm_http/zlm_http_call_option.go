package zlm_http

import (
	"net/http"
)

// CallSettings allow fine-grained control over how calls are made.
type CallSettings struct {
	// Content-Type
	contentType string
	// Accept
	accept string
	// custom header
	header http.Header
	// baseUrl overwrite api call
	baseUrl string
	// path overwrite api call
	path string
	// 是否是文件下载
	isFileDownload bool
}

// CallOption is an option used by Invoke to control behaviors of RPC calls.
// CallOption works by modifying relevant fields of CallSettings.
type CallOption func(cs *CallSettings)

// WithCoContentType use encoding.MIMExxx
func WithCoContentType(contentType string) CallOption {
	return func(cs *CallSettings) {
		cs.contentType = contentType
	}
}

// WithCoAccept use encoding.MIMExxx
func WithCoAccept(accept string) CallOption {
	return func(cs *CallSettings) {
		cs.accept = accept
	}
}

// WithCoBaseUrl
func WithCoBaseUrl(baseUrl string) CallOption {
	return func(cs *CallSettings) {
		cs.baseUrl = baseUrl
	}
}

// WithCoPath
func WithCoPath(path string) CallOption {
	return func(cs *CallSettings) {
		cs.path = path
	}
}

// WithCoHeader
func WithCoHeader(k, v string) CallOption {
	return func(cs *CallSettings) {
		cs.header.Add(k, v)
	}
}

// WithFileDownload
func WithFileDownload() CallOption {
	return func(cs *CallSettings) {
		cs.isFileDownload = true
	}
}
