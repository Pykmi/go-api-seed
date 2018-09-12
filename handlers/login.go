package handlers

import (
	"fmt"
	"net/http"

	"bitbucket.org/pykmiteam/mock-api/datastore"
	"bitbucket.org/pykmiteam/mock-api/logger"
)

// LoginRequest : one potato, two potato...
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

const loginUser = "admin"
const loginPass = "admin"

// Login : Handler for login requests
func Login(w http.ResponseWriter, r *http.Request) {
	data := LoginRequest{}
	data.Username = r.URL.Query().Get("user")
	data.Password = r.URL.Query().Get("pass")

	status := http.StatusOK

	store := datastore.Get(r)
	auth, err := store.Authenticate(data.Username, data.Password)
	if err != nil {
		fmt.Println(err)
	}

	if auth == false {
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
