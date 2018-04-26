package handlers

import (
	"net/http"
)

const loginUser = "admin"
const loginPass = "admin"

func Login(w http.ResponseWriter, r *http.Request) {
	data := LoginRequest{}
	data.Username = r.URL.Query().Get("user")
	data.Password = r.URL.Query().Get("pass")

	status := http.StatusOK

	if data.Username != loginUser || data.Password != loginPass {
		status = http.StatusUnauthorized
	}

	// print response to the http writer
	w.WriteHeader(status)
	w.Write([]byte(""))
}
