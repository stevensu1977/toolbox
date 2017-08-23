package mgo

import "reflect"

type MgoboxDocument struct {
	Model      reflect.Type
	Collection string
	BSONTag    []string
}
