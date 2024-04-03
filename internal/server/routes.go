package server

import (
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("createuser", CreateUserHandler)
}
