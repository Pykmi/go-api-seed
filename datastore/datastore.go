package datastore

import (
	"context"
	"net/http"

	"gopkg.in/mgo.v2"
)

// StoreOptions : options struct for the datastore
type StoreOptions struct {
	Host     	string
	Port			string
	User			string
	Pass			string
	Database	string
}

// toString : return url string
func (opt *StoreOptions) toString() string {
	output := opt.User + ":" + opt.Pass
	output += "@" + opt.Host + ":" + opt.Port
  return "mongodb://" + output
}

// Store : the datastore struct
type Store struct {
	cli				*mgo.Session
	host      string
	database	string
}

// the key type is unexported to avoid context collision
type key int

const contextKey key = iota

// Key : returns the contextKey, only for testing purposes
func Key() key {
	return contextKey
}

// Get : returns the store object from http context
func Get(r *http.Request) *Store {
	return r.Context().Value(contextKey).(*Store)
}

// Middleware : stores the datastore into the http context
func Middleware(store *Store) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(
			w http.ResponseWriter,
			r *http.Request,
		) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, contextKey, store)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// New : creates a database connection
func New(opt StoreOptions) *Store {
	store := &Store{}

	session, err := mgo.Dial(opt.toString())
	if err != nil {
		panic(err)
	}

	store.cli = session
	store.database = opt.Database
	return store
}
