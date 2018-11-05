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

// AuthByEmail : authenticate user by email address
func (s *Store) AuthByEmail(email string, pass string) (bool, error) {
	c := s.cli.DB(s.database).C(users)

	Usr := User{}
	query := bson.M{ "email": email, "password": pass };

	err := c.Find(query).One(&Usr)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *Store) getUserByEmail(email string) (User, error) {
	c := s.cli.DB(s.database).C(users)
	result := User{}
	err := c.Find(bson.M{"email": email}).One(&result)
	if err != nil {
		return User{}, err
	}

	return result, nil;
}
