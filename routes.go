package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

	"bitbucket.org/pykmiteam/mock-api/handlers"
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

func setupRoutes() http.Handler {
	// API Routes
	router := goji.NewMux()

	// add cors
	router.Use(corsMiddle)

	router.HandleFunc(pat.Get("/"), handlers.Default)
	router.HandleFunc(pat.Post("/login"), handlers.Login)

	return router
}
