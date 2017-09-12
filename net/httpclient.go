package net

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/stevensu1977/toolbox/storage"
)

type HTTPSimpleClient struct {
	http.Client
	Headers map[string]string
}

type HTTPSimpleResult struct {
	StatusCode int
	Body       []byte
}

var headers = map[string]string{
	"User-Agent":   "curl/7.43.0",
	"Accept":       "*/*",
	"Content-Type": "json",
}

func setHeader(req *http.Request, headers map[string]string) {
	for k, v := range headers {
		req.Header.Set(k, v)
	}

}

// NewHTTPSimpleClient provide init httpsimpleclient
func NewHTTPSimpleClient() *HTTPSimpleClient {
	return &HTTPSimpleClient{
		Headers: map[string]string{},
	}
}

// Get 方法提供快速访问URL
func (c *HTTPSimpleClient) Get(url string) (HTTPSimpleResult, error) {
	return c.DoSimpleRequest(http.MethodGet, url, nil)
}

// Post 方法提供快速访问URL
func (c *HTTPSimpleClient) Post(url string, data interface{}) (HTTPSimpleResult, error) {

	return c.DoSimpleRequest(http.MethodPost, url, data)
}

// Put 方法提供快速访问URL
func (c *HTTPSimpleClient) Put(url string, data interface{}) (HTTPSimpleResult, error) {
	return c.DoSimpleRequest(http.MethodPut, url, data)
}

// Delete 方法提供快速访问URL
func (c *HTTPSimpleClient) Delete(url string, data interface{}) (HTTPSimpleResult, error) {
	return c.DoSimpleRequest(http.MethodDelete, url, data)
}

// PostFormData 方法提供通过Form表单快速访问URL
func (c *HTTPSimpleClient) PostFormData(url string, data url.Values) (HTTPSimpleResult, error) {
	result := HTTPSimpleResult{StatusCode: -1, Body: nil}
	resp, err := c.PostForm(url, data)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	result.StatusCode = resp.StatusCode

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	result.Body = body
	return result, nil
}

// DoSimpleRequest 封装了所有Method的方法提供快速访问
func (c *HTTPSimpleClient) DoSimpleRequest(method, url string, data interface{}) (HTTPSimpleResult, error) {

	result := HTTPSimpleResult{StatusCode: -1, Body: nil}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return result, err
	}

	if data != nil {
		dataByte, err := json.Marshal(data)

		if err != nil {
			return result, err
		}
		body := bytes.NewBuffer(dataByte)
		rc := ioutil.NopCloser(body)

		req.Body = rc
	}

	if c.Headers == nil {
		setHeader(req, headers)
	} else {
		setHeader(req, c.Headers)
	}

	rep, err := c.Do(req)
	if err != nil {
		return result, err
	}

	defer rep.Body.Close()
	body, err := ioutil.ReadAll(rep.Body)
	result.StatusCode = rep.StatusCode
	if err != nil {

		return result, err
	}

	result.Body = body
	return result, nil

}

// GetUnmarshal 工具类, Get 并解析返回的报文，返回 error
func GetUnmarshal(url string, ret interface{}) (err error) {
	log.Printf("url=%s \n", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}

	return nil
}

// PostMarshal 工具类, POST 编组并返回 error
func PostMarshal(url string, v interface{}) (err error) {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}

	err = PostUnmarshal(url, js, nil)
	if err != nil {
		return err
	}

	return nil
}

// Post 工具类, POST json 并返回 error
func Post(url string, js []byte) (err error) {

	err = PostUnmarshal(url, js, nil)
	if err != nil {
		return err
	}

	return nil
}

// PostMarshalUnmarshal 工具类, POST 编组并解析返回的报文，返回 error
func PostMarshalUnmarshal(url string, v interface{}, ret interface{}) (err error) {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return PostUnmarshal(url, js, ret)
}

// PostUnmarshal 工具类, POST json 并解析返回的报文，返回 error
func PostUnmarshal(url string, js []byte, ret interface{}) (err error) {
	log.Printf("url=%s, body=%s", url, js)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(js))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}

	return nil
}

// Upload 工具类, 上传文件
func Upload(url, fieldName string, file *os.File, ret interface{}, desc ...string) (err error) {
	buf := &bytes.Buffer{}
	w := multipart.NewWriter(buf)

	//关键的一步操作
	fw, err := w.CreateFormFile(fieldName, file.Name())
	if err != nil {
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		return err
	}
	contentType := w.FormDataContentType()
	if len(desc) > 0 {
		w.WriteField("description", desc[0])
	}
	w.Close()

	log.Printf("url=%s, fieldName=%s, fileName=%s", url, fieldName, file.Name())
	resp, err := http.Post(url, contentType, buf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}

	return nil
}

func Fetch(path string, root ...string) error {
	urlPtr, err := url.Parse(path)
	rootPath := "."
	if len(root) > 0 {
		rootPath = root[0]
	}
	if err != nil {
		return err
	}
	parentPath := filepath.Dir(urlPtr.Host + urlPtr.Path)
	fileName := filepath.Base(urlPtr.Path)
	if strings.HasSuffix(path, "/") {
		fileName = "index.html"
	}

	resp, err := http.Get(path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = storage.MkdirAll(rootPath + "/" + parentPath)
	if err != nil {
		return err
	}
	w, err := os.Create(rootPath + "/" + parentPath + "/" + fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
