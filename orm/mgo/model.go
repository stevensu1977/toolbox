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

func (self *MgoboxDocument) Count(query ...bson.M) (int, error) {
	checkInit()
	mConn := Conn()
	defer mConn.Close()
	if len(query) > 0 {
		return mConn.DB(Database()).C(self.Collection).Find(query[0]).Select(bson.M{"_id": 1}).Count()
	}
	return mConn.DB(Database()).C(self.Collection).Find(nil).Select(bson.M{"_id": 1}).Count()
}

func (self *MgoboxDocument) Page(query ...bson.M) (int, error) {

	checkInit()
	mConn := Conn()
	defer mConn.Close()
	count := 1
	if len(query) > 0 {
		count, _ = self.Count(query[0])
	} else {
		count, _ = self.Count()
	}

	if count < pageSize {
		return 1, nil
	}

	odd := count % pageSize
	page := (count - odd) / pageSize

	if odd > 0 {
		page = page + 1
	}

	return page, nil

}

func (self *MgoboxDocument) Find(query bson.M) *mgo.Query {
	checkInit()
	if query == nil {
		return session.DB(Database()).C(self.Collection).Find(nil)
	}
	return session.DB(Database()).C(self.Collection).Find(query)

}

func (self *MgoboxDocument) FindWithPage(page int, query bson.M) *mgo.Query {
	checkInit()

	count, err := self.Count(query)
	if err != nil {
		panic(err)
	}
	if page < 1 {
		page = 1
	}

	if count < pageSize*(page-1) {
		page = 2
	}

	if query == nil {
		return session.DB(Database()).C(self.Collection).Find(nil).Skip(pageSize * (page - 1)).Limit(pageSize)
	}
	return session.DB(Database()).C(self.Collection).Find(query).Skip(pageSize * (page - 1)).Limit(pageSize)

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
