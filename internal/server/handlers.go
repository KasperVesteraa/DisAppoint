package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/KasperVesteraa/DisAppoint/internal/api"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to DisAppoint!\n")
}

// func CreateUserHandler(db *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var user api.User
// 		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 			http.Error(w, err.Error(), http.StatusBadRequest)
// 			return
// 		}
// 		_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusCreated)
// 		fmt.Fprintf(w, "User created successfully!\n")
// 	}
// }

func UserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var user api.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "User created successfully!\n")
		default:
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	}
}
