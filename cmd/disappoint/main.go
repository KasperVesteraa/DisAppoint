package main

import (
	"fmt"
	"net/http"

	"github.com/KasperVesteraa/DisAppoint/internal/server"
)

func main() {

	server.InitializeRoutes()
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
