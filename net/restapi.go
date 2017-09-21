package net

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

const (
	API_JSON = iota //API_JSON
	API_XML
	API_RAW
)

//Encoder 抽象出来的处理器
type Encoder interface {
	Encode(v interface{}) error
}

type XMLEncoder struct {
	encoder *xml.Encoder
}

type JSONEncoder struct {
	encoder *json.Encoder
}

type RawEncoder struct {
	body *bytes.Buffer
}

func (r *RawEncoder) Encode(v interface{}) error {
	_, err := r.body.Write([]byte(fmt.Sprintf("%v", v)))
	return err
}

func (j *JSONEncoder) Encode(v interface{}) error {
	return j.encoder.Encode(v)
}
func (x *XMLEncoder) Encode(v interface{}) error {
	return x.encoder.Encode(v)
}

//ServerJSON provide simple REST API JSON write
func ServerJSON(w http.ResponseWriter, model interface{}) {
	serverAPI(API_JSON, w, model)
}

func ServerXML(w http.ResponseWriter, model interface{}) {
	serverAPI(API_XML, w, model)
}

func ServerRAW(w http.ResponseWriter, model interface{}) {
	serverAPI(API_RAW, w, model)
}
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
