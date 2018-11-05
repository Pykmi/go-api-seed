package datastore

import (
	"testing"

	/* "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson" */
)

const host string = "0.0.0.0"
const port string = "8081"
const user string = "pykmi"
const pass string = "okilzw"

const mongoDB string = "pykmi-dev-db"
const mongoC string = "users"

const adminEmail string = "tomi.kaistila@tieto.com"
const connect string = "mongodb://pykmi:okilzw@0.0.0.0:8081"

func TestNew(t *testing.T) {
	opt := StoreOptions{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Database: mongoDB,
	}

	store := New(opt)

	/* c := session.DB(mongoDB).C(mongoC)

	result := User{}
	err = c.Find(bson.M{"username": "pykmi"}).One(&result)
	if err != nil {
		t.Fatal(err)
	} */

	data, err := store.getUser(adminEmail)
	if err != nil {
		t.Fatal(err)
	}

	if data.Email != adminEmail {
		t.Fatalf("\nExpected v=%v got=%v", adminEmail, data.Email)
	}
}
