package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"

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

func setupRoutes(EventLogger *logger.Logger) http.Handler {
	// API Routes
	router := goji.NewMux()

	// add cors
	router.Use(corsMiddle)
	router.Use(logger.Middleware(EventLogger))

	router.HandleFunc(pat.Get("/"), handlers.Default)
	router.HandleFunc(pat.Get("/api/login"), handlers.Login)

	return router
}
