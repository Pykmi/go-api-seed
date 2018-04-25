package handlers

import (
	"encoding/json"
	"net/http"
)

const LoginUser = "admin"
const LoginPass = "admin"

func Login(w http.ResponseWriter, r *http.Request) {
	var data LoginRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	res := Response{
		Code:   Code[200],
		Status: Status[200],
	}

	if err != nil {
		res.Code = Code[500]
		res.Status = Status[500]
		res.Error = err.Error()
	}

	if data.Username != LoginUser || data.Password != LoginPass {
		res.Code = Code[401]
		res.Status = Status[401]
	}

	// print response to the http writer
	send := writeJSON(w, res)
	httpError(w, send)
}
