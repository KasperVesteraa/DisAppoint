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

func UserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost: // Create
			var user api.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "User created successfully!\n")
		case http.MethodDelete: // Delete
			var user api.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			_, err := db.Exec("DELETE FROM users WHERE email = $1", user.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "User deleted successfully!\n")
		case http.MethodPut: // Update
			var user api.User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			_, err := db.Exec("UPDATE users SET name = $1, password = $2 WHERE email = $3", user.Name, user.Password, user.Email)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "User updated succesfully!\n")
		default:
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	}
}
