package mgo

import (
	"errors"
	"fmt"
	"log"
	"reflect"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session

var models = make(map[string]*MgoboxDocument)

var _mongoURL = "mongodb://localhost:27017/apicatalog"

var _database = "mgobox"

var _init = false

func init() {

}

func checkInit() {
	if _init == false {
		panic(errors.New("Context need  init first"))
	}
}
func Conn() *mgo.Session {
	return session.Copy()
}

func Close() {
	session.Close()
}

func PrintModels() {
	fmt.Println(models)
}

func InitContext(mongoURL, database string) {
	_mongoURL = mongoURL
	_database = database
	_init = true
	log.Printf("MongoDB url [%s] ", _mongoURL)
	sess, err := mgo.Dial(_mongoURL)
	if err != nil {
		panic(err)
	}
	session = sess
	session.SetMode(mgo.Monotonic, true)
}

func DefaultInit() {
	InitContext(_mongoURL, _database)
}

func SetDatabase(database string) {
	_database = database
}

func Database() string {
	return _database
}

func Model(model interface{}) (*MgoboxDocument, error) {

	v := reflect.ValueOf(model)
	i := reflect.Indirect(v)
	s := i.Type()
	if _, ok := models[s.String()]; ok {
		return models[s.String()], nil
	}
	return nil, errors.New("Context not fond " + s.String() + " model, please register first")
}

func ModelByString(packageName, modelName string) (*MgoboxDocument, error) {

	if _, ok := models[packageName+"."+modelName]; ok {
		return models[packageName+"."+modelName], nil
	}

	return nil, errors.New("Context not fond " + packageName + "." + modelName + " model, please register first")

}

func Register(model interface{}, collectionName string) error {

	v := reflect.ValueOf(model)
	i := reflect.Indirect(v)
	s := i.Type()

	log.Printf("[Register] Model: %s, Collection: %s", s.String(), collectionName)

	tags := make([]string, 0)

	haveID := false

	for i := 0; i < s.NumField(); i++ {
		bsonTag := s.Field(i).Tag.Get("bson")
		if bsonTag != "" {
			if bsonTag == "_id" {
				haveID = true
			}
			tags = append(tags, s.Field(i).Tag.Get("bson"))
		}

	}
	if len(tags) == 0 || haveID == false {
		return errors.New("Model not have bson tag/ Primary Key(etc._id) !")
	}

	m := &MgoboxDocument{}
	m.Model = s
	m.BSONTag = tags
	m.Collection = collectionName

	models[s.String()] = m

	return nil
}
