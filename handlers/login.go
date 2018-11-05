package handlers

import (
	"fmt"
	"net/http"

	"github.com/pykmi/go-api-seed/datastore"
	"github.com/pykmi/go-api-seed/logger"
)

// LoginRequest : one potato, two potato...
type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

// Login : Handler for login requests
func Login(w http.ResponseWriter, r *http.Request) {
	data := LoginRequest{}
	data.Email = r.URL.Query().Get("email")
	data.Password = r.URL.Query().Get("pass")

	status := http.StatusOK

	store := datastore.Get(r)
	auth, err := store.AuthByEmail(data.Email, data.Password)
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
