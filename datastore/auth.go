package datastore

import (
	as "github.com/aerospike/aerospike-client-go"
)

const bucket string = "Users"

// Authenticate : authenticate users
func (s *Store) Authenticate(user string, pass string) (bool, error) {
	key, err := as.NewKey("test", bucket, user)
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

	return true, nil
}
