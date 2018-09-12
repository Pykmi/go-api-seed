package datastore

import (
	"context"
	"net/http"

	as "github.com/aerospike/aerospike-client-go"
)

// StoreOptions : options struct for the datastore
type StoreOptions struct {
	Bucket    string
	Host      string
	Namespace string
	Port      int
}

// Store : the datastore struct
type Store struct {
	cli       *as.Client
	host      string
	namespace string
	port      int
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

	cli, err := as.NewClient(opt.Host, opt.Port)
	if err != nil {
		panic(err)
	}

	store.cli = cli
	return store
}
