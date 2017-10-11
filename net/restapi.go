package net

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	//API_JSON json type
	API_JSON = iota
	API_XML
	API_RAW
)

//Encoder 抽象出来的Encoder处理器
type Encoder interface {
	Encode(v interface{}) error
}

//Decoder 抽象出来的Decoder处理器
type Decoder interface {
	Decode(v interface{}) error
}

//XMLDecoder xml decorder
type XMLDecoder struct {
	decoder *xml.Decoder
}

//JSONDecoder json decorder
type JSONDecoder struct {
	decoder *json.Decoder
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

//Decode implement Decoder interface
func (j *JSONDecoder) Decode(v interface{}) error {
	return j.decoder.Decode(v)
}

//Decode implement Decoder interface
func (x *XMLDecoder) Decode(v interface{}) error {
	return x.decoder.Decode(v)
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

//ConsumeJSON provid simple JSON decode from request body
func ConsumeJSON(r io.ReadCloser, model interface{}) {
	consumeAPI(API_JSON, r, model)
}

//ConsumeXML provid simple XML decode from request body
func ConsumeXML(r io.ReadCloser, model interface{}) {
	consumeAPI(API_XML, r, model)
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

//internal func
func consumeAPI(apiType int, r io.ReadCloser, model interface{}) error {
	var contentWriter Decoder
	switch apiType {
	case API_JSON:
		contentWriter = &JSONDecoder{
			decoder: json.NewDecoder(r),
		}
	case API_XML:
		contentWriter = &XMLDecoder{
			decoder: xml.NewDecoder(r),
		}
	default:
		return errors.New("only support json, xml decode")

	}

	err := contentWriter.Decode(model)
	if err != nil {
		return err
	}
	return nil
}
