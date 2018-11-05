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

func TestNew(t *testing.T) {
	opt := StoreOptions{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Database: mongoDB,
	}

	store := New(opt)

	data, err := store.getUserByEmail(adminEmail)
	if err != nil {
		t.Fatal(err)
	}

	if data.Email != adminEmail {
		t.Fatalf("\nExpected v=%v got=%v", adminEmail, data.Email)
	}
}

func TestAuthByEmail(t *testing.T) {
	opt := StoreOptions{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Database: mongoDB,
	}

	store := New(opt)
	data, err := store.AuthByEmail(adminEmail, "admin")
	if err != nil {
		t.Fatal(err)
	}

	if data.Email != adminEmail {
		t.Fatalf("\nExpected v=%v got=%v", adminEmail, data.Email)
	}
}
