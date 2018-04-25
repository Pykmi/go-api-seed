package handlers

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Error  string `json:"error,omitempty"`
	Code   int    `json:"code"`
	Status string `json:"status"`
}

var (
	Status = map[int]string{
		200: "OK",
		400: "Bad Request",
		401: "Unauthorized",
		404: "Not Found",
		500: "Internal Server Error",
	}
	Code = map[int]int{
		200: 200,
		400: 400,
		401: 401,
		404: 404,
		500: 500,
	}
)
