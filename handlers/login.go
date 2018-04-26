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
	res := response{}
	res.code = Code[200]
	res.status = Status[200]

	if err != nil {
		res.code = Code[500]
		res.status = Status[500]
	}

	if data.Username != LoginUser || data.Password != LoginPass {
		res.code = Code[401]
		res.status = Status[401]
	}

	res.status = joinStatusCode(res.code, res.status)
	res.request = r

	// print response to the http writer
	httpWrite(res)
}
