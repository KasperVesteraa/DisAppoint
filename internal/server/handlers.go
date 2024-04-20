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
			row := db.QueryRow("SELECT id FROM users WHERE email = $1", user.Email)
			var id string
			if err := row.Scan(&id); err == nil {
				http.Error(w, "User already exists", http.StatusConflict)
				return
			}
			// user.CreateUuid()
			_, err := db.Exec("INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)", user.Id, user.Name, user.Email, user.Password)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "User created successfully!\n")

		case http.MethodGet: // Read
			email := r.URL.Query().Get("email")
			if email == "" {
				http.Error(w, "Email parameter required", http.StatusBadRequest)
				return
			}
			var user api.User
			row := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email)
			err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(user)

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
		default:
			http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
		}
	}
}

func AppointmentHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost: // Create
			var appt api.Appointment
			if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			_, err := db.Exec("INSERT INTO appointments (id, title, location, description, start_time, end_time, author_id) VALUES ($1, $2, $3, $4, $5, $6, $7)", appt.Id, appt.Title, appt.Location, appt.Description, appt.StartTime, appt.EndTime, appt.AuthorId)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			for _, part_id := range appt.Parts_id {
				_, err := db.Exec("INSERT INTO appointment_participants (appointment_id, user_id) VALUES ($1, $2)", appt.Id, part_id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			w.WriteHeader(http.StatusCreated)
			fmt.Fprintf(w, "Appointment created successfully!\n")

		case http.MethodGet: // Read
		case http.MethodPut: // Update
		case http.MethodDelete: // Delete
		}
	}
}

// type Appointment struct {
// 	Id           string
// 	Title        string
// 	Location     string
// 	Description  string
// 	StartTime    int
// 	EndTime      int
// 	Author       User
// 	Participants []User
// }

// CREATE TABLE IF NOT EXISTS users (
//     id VARCHAR(255) PRIMARY KEY,
//     name VARCHAR(255) NOT NULL,
//     email VARCHAR(255) UNIQUE NOT NULL,
//     password TEXT NOT NULL
// );

// CREATE TABLE IF NOT EXISTS appointments (
//     id VARCHAR(255) PRIMARY KEY,
//     title VARCHAR(255) NOT NULL,
//     location TEXT,
//     description TEXT,
//     start_time INT NOT NULL,
//     end_time INT NOT NULL,
//     author_id VARCHAR(255),
//     FOREIGN KEY (author_id) REFERENCES users (id)
// );

// CREATE TABLE IF NOT EXISTS appointment_participants (
//     appointment_id VARCHAR(255),
//     user_id VARCHAR(255),
//     PRIMARY KEY (appointment_id, user_id),
//     FOREIGN KEY (appointment_id) REFERENCES appointments (id) ON DELETE CASCADE,
//     FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
// );
