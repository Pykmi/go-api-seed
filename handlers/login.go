package handlers

import (
	"net/http"

	"bitbucket.org/pykmiteam/mock-api/logger"
)

const loginUser = "admin"
const loginPass = "admin"

// Login : Handler for login requests
func Login(w http.ResponseWriter, r *http.Request) {
	data := LoginRequest{}
	data.Username = r.URL.Query().Get("user")
	data.Password = r.URL.Query().Get("pass")

	status := http.StatusOK

	if data.Username != loginUser || data.Password != loginPass {
		status = http.StatusUnauthorized
	}

	e := logger.NewEvent()
	e.RemoteAddr = r.RemoteAddr
	e.Method = "GET"
	e.Status = status
	e.RequestURI = r.RequestURI

	logger := logger.Get(r)
	logger.Log(e)

	// print response to the http writer
	w.WriteHeader(status)
	w.Write([]byte(""))
}
