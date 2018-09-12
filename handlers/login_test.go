package handlers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/pykmiteam/mock-api/datastore"
	"bitbucket.org/pykmiteam/mock-api/logger"
)

const dbHost string = "127.0.0.1"
const dbPort int = 9000

func TestLoginHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/login?user=admin&pass=admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Login)

	// create datastore
	StoreOpt := datastore.StoreOptions{Host: dbHost, Namespace: "test", Port: dbPort}
	Store := datastore.New(StoreOpt)

	// create event logger
	EventLogger := logger.New()

	// save datastore and event logger to the context
	ctx := req.Context()
	ctx = context.WithValue(ctx, datastore.Key(), Store)
	ctx = context.WithValue(ctx, logger.Key(), EventLogger)

	req = req.WithContext(ctx)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v wanted %v",
			status, http.StatusOK)
	}
}
