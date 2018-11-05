package datastore

import (
	"gopkg.in/mgo.v2/bson"
)

// User : these fields must be found on the collection
type User struct {
	Username string
	Password string
	Name string
	Email string
}

const users string = "users"

// Authenticate : authenticate users
func (s *Store) Authenticate(user string, pass string) (bool, error) {
	/* key, err := as.NewKey("test", bucket, user)
	if err != nil {
		return false, err
	}

	rec, err := s.cli.Get(nil, key)
	if err != nil {
		return false, err
	}

	if rec == nil {
		return false, nil
	}

	if rec.Bins["password"] != pass {
		return false, nil
	}

	return true, nil */
	return false, nil
}

func (s *Store) getUser(email string) (User, error) {
	c := s.cli.DB(s.database).C(users)
	result := User{}
	err := c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		return User{}, err
	}

	return result, nil;
}
