package net

import (
	"encoding/json"
	"log"
	"net/http"
)

type HTTPSimpleREST struct {
	HTTPSimpleClient
}

// NewHTTPSimpleREST provide init httpsimpleclient
func NewHTTPSimpleREST() *HTTPSimpleREST {

	c := &HTTPSimpleREST{}
	c.Headers = map[string]string{}
	return c
}

// Get 提供REST方法
func (c *HTTPSimpleREST) Get(url string, response interface{}) error {
	return c.DoREST(http.MethodGet, url, nil, response)
}

// Post 提供REST方法
func (c *HTTPSimpleREST) Post(url string, data interface{}, reponse ...interface{}) error {
	switch len(reponse) == 0 {
	case false:
		return c.DoREST(http.MethodPost, url, data, reponse[0])
	default:
		return c.DoREST(http.MethodPost, url, data, nil)
	}
}

// Put 提供REST方法
func (c *HTTPSimpleREST) Put(url string, data interface{}, reponse ...interface{}) error {
	switch len(reponse) == 0 {
	case false:
		return c.DoREST(http.MethodPut, url, data, reponse[0])
	default:
		return c.DoREST(http.MethodPut, url, data, nil)
	}
}

// Delete 提供REST方法
func (c *HTTPSimpleREST) Delete(url string, data interface{}, reponse ...interface{}) error {
	switch len(reponse) == 0 {
	case false:
		return c.DoREST(http.MethodDelete, url, data, reponse[0])
	default:
		return c.DoREST(http.MethodDelete, url, data, nil)
	}
}

// DoREST 提供REST方法统一封装
func (c *HTTPSimpleREST) DoREST(method, url string, data interface{}, ret interface{}) error {

	result, err := c.DoSimpleRequest(method, url, data)
	if err != nil {
		return err
	}

	switch ret == nil {
	case true:
		log.Printf("Debug: %d %s", result.StatusCode, result.Body)
	case false:
		err = json.Unmarshal(result.Body, ret)
		if err != nil {
			return err
		}

	}

	return nil

}
