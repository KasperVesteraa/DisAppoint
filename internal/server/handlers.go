package server

import (
	"fmt"
	"net/http"

	"github.com/KasperVesteraa/DisAppoint/internal/api"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to DisAppoint!\n")
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CreateUser entered...!\n")
	api.CreateUser("1", "Kasper", "kasper@onlyfans.com", "1234")
}
