package mgo

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (self *MgoboxDocument) New(model interface{}) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).Insert(model)

}

func (self *MgoboxDocument) Find(query bson.M) *mgo.Query {
	checkInit()
	if query == nil {
		return session.DB(Database()).C(self.Collection).Find(nil)
	}
	return session.DB(Database()).C(self.Collection).Find(query)

}

func (self *MgoboxDocument) FindId(id string) *mgo.Query {
	checkInit()
	return session.DB(Database()).C(self.Collection).FindId(bson.ObjectIdHex(id))
}

func (self *MgoboxDocument) Remove(query bson.M) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).Remove(query)
}

func (self *MgoboxDocument) RemoveId(id string) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).RemoveId(bson.ObjectIdHex(id))
}

func (self *MgoboxDocument) UpdateId(id string, update interface{}) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).UpdateId(bson.ObjectIdHex(id), update)
}

func (self *MgoboxDocument) Update(query bson.M, update interface{}) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).Update(query, update)
}

func (self *MgoboxDocument) Change(query bson.M, change bson.M) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).Update(query, bson.M{"$set": change})
}

func (self *MgoboxDocument) ChangeId(id string, change interface{}) error {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	return mConn.DB(Database()).C(self.Collection).UpdateId(bson.ObjectIdHex(id), bson.M{"$set": change})
}

func ChangeModel(input, empty interface{}) {

}
