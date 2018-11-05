package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/pykmi/go-api-seed/datastore"
	"github.com/pykmi/go-api-seed/logger"
)

const dbHost string = "127.0.0.1"
const dbPort int = 9000

const host string = "0.0.0.0"
const port string = "8081"
const user string = "pykmi"
const pass string = "okilzw"

const mongoDB string = "pykmi-dev-db"
const mongoC string = "users"

func TestLoginHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/login?email=tomi.kaistila@tieto.com&pass=admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	// create datastore
	opt := db.StoreOptions{
		Host: host,
		Port: port,
		User: user,
		Pass: pass,
		Database: mongoDB,
	}

	store := db.New(opt)

	// create event logger
	EventLogger := logger.New() 

	// save datastore and event logger to the context
	ctx := req.Context()
	ctx = context.WithValue(ctx, db.Key(), store)
	ctx = context.WithValue(ctx, logger.Key(), EventLogger)

	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}
}
