package mgo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `bson:"_id"      json:"id,omitempty"`
	Email    string        ` bson:"email" json:"email,omitempty"`
	Password string        ` bson:"password" json:"password,omitempty"`
}

type Spec struct {
	ID      bson.ObjectId `bson:"_id"      json:"id,omitempty"`
	Version string        ` bson:"version" json:"version,omitempty"`
	App     string        ` bson:"pass" json:"app,omitempty"`
}

func TestRegisterSuccess(t *testing.T) {
	DefaultInit()
	err := Register(&User{}, "users")
	if err != nil {
		t.Error(err)
	}

	_, err = Model(&User{})

	if err != nil {
		t.Error(err)
	}

	_, err = ModelByString("mgobox", "User")

	if err != nil {
		t.Error(err)
	}
}

func TestCountAndPage(t *testing.T) {
	DefaultInit()
	err := Register(&User{}, "users")
	if err != nil {
		t.Error(err)
	}

	userModel, err := Model(&User{})
	users := &[]User{}
	PageSize(1)
	t.Log(userModel.Count())
	t.Log(userModel.Page())
	userModel.FindWithPage(2, nil).All(users)
	t.Log(PageSize(), users)
	PageSize(10)

	t.Log(userModel.Page())
	userModel.FindWithPage(2, nil).All(users)
	t.Log(PageSize(), users)
}

func TestCRUD(t *testing.T) {
	DefaultInit()

	email := "wangzheng@orientsoft.cn"

	err := Register(&User{}, "users")
	if err != nil {
		t.Error(err)
	}

	userModel, err := Model(&User{})

	if err != nil {
		t.Error(err)
	}

	user := &User{}
	userModel.Find(bson.M{"email": email}).One(user)

	if user.ID.Hex() != "" {
		t.Log("User find ")
	} else {
		user.Email = "wangzheng@orientsoft.cn"
		user.Password = "welcome1"
		user.ID = bson.NewObjectId()
		err = userModel.New(user)
		if err != nil {
			t.Error(err)
		} else {
			t.Log("New user success")
		}
	}
	users := &[]User{}
	userModel.Find(nil).All(users)
	t.Log(users)
	userModel.RemoveId(user.ID.Hex())
}
