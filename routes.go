package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"bitbucket.org/pykmiteam/mock-api/datastore"
	"bitbucket.org/pykmiteam/mock-api/handlers"
	"bitbucket.org/pykmiteam/mock-api/logger"
)

// corsMiddle middleware handles the supported access-control headers
func corsMiddle(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS, DELETE")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func setupRoutes(EventLogger *logger.Logger, Store *datastore.Store) http.Handler {
	// create router
	router := goji.NewMux()

	// add cors
	router.Use(corsMiddle)
	router.Use(logger.Middleware(EventLogger))
	router.Use(datastore.Middleware(Store))

	// api routes
	router.HandleFunc(pat.Get("/api/login"), handlers.Login)

	return router
}
