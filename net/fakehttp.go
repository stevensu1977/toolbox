package net

import (
	"bytes"
	"net/http"
)

type FakeResponseWriter struct {
	headers http.Header
	body    bytes.Buffer
	status  int
}

func NewFakeResponseWriter() *FakeResponseWriter {

	return &FakeResponseWriter{
		headers: http.Header{},
	}

}

func (r *FakeResponseWriter) Header() http.Header {
	return r.headers
}

func (r *FakeResponseWriter) WriteHeader(status int) {
	r.status = status
}

func (r *FakeResponseWriter) Write(body []byte) (int, error) {
	return r.body.Write(body)
}

func (r *FakeResponseWriter) Body() []byte {
	return r.body.Bytes()
}
