package net

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	//API_JSON json type
	API_JSON = iota
	API_XML
	API_RAW
)

//Encoder 抽象出来的处理器
type Encoder interface {
	Encode(v interface{}) error
}

//XMLEncoder xml encoder
type XMLEncoder struct {
	encoder *xml.Encoder
}

//JSONEncoder json struct
type JSONEncoder struct {
	encoder *json.Encoder
}

//RawEncoder json struct
type RawEncoder struct {
	body *bytes.Buffer
}

//Encode implement Encoder interface
func (r *RawEncoder) Encode(v interface{}) error {
	_, err := r.body.Write([]byte(fmt.Sprintf("%v", v)))
	return err
}

//Encode implement Encoder interface
func (j *JSONEncoder) Encode(v interface{}) error {
	return j.encoder.Encode(v)
}

//Encode implement Encoder interface
func (x *XMLEncoder) Encode(v interface{}) error {
	return x.encoder.Encode(v)
}

//ServerJSON provide simple REST API go struct to JSON  func
func ServerJSON(w http.ResponseWriter, model interface{}) {
	serverAPI(API_JSON, w, model)
}

//ServerXML provide simple REST API go struct to XML func
func ServerXML(w http.ResponseWriter, model interface{}) {
	serverAPI(API_XML, w, model)
}

//ServerRAW provid simple REST API RAW go structure
func ServerRAW(w http.ResponseWriter, model interface{}) {
	serverAPI(API_RAW, w, model)
}

//serverAPI is internal func
func serverAPI(apiType int, w http.ResponseWriter, model interface{}) {
	var buf bytes.Buffer
	var contentWriter Encoder
	var contentType string
	switch apiType {
	case API_JSON:
		contentWriter = &JSONEncoder{
			encoder: json.NewEncoder(&buf),
		}
		contentType = "application/json"
	case API_XML:
		contentWriter = &XMLEncoder{
			encoder: xml.NewEncoder(&buf),
		}
		contentType = "text/xml"
	default:
		contentWriter = &RawEncoder{
			body: &buf,
		}
		contentType = "text/html"
	}

	err := contentWriter.Encode(model)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("Content-Type", contentType)
	buf.WriteTo(w)
}
