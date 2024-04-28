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
			row := db.QueryRow("SELECT id FROM users WHERE email = $1", user.Email) // Check if user already exists
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

		case http.MethodDelete: // Delete
			author_id := r.URL.Query().Get("author_id")
			if author_id == "" {
				http.Error(w, "Author id parameter required", http.StatusBadRequest)
				return
			}
			start_time := r.URL.Query().Get("start_time")
			if start_time == "" {
				http.Error(w, "Start time parameter required", http.StatusBadRequest)
				return
			}
			res, err := db.Exec("DELETE FROM appointments WHERE author_id = $1 AND start_time = $2", author_id, start_time)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			if rows, _ := res.RowsAffected(); rows == 0 {
				http.Error(w, "No appointment found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Appointment deleted successfully!\n")

		case http.MethodPost: // Create
			var appt api.Appointment
			if err := json.NewDecoder(r.Body).Decode(&appt); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()
			row := db.QueryRow("SELECT title FROM appointments WHERE author_id = $1 AND start_time = $2", appt.AuthorId, appt.StartTime) // Check if appointment already exists
			var title string
			if err := row.Scan(&title); err == nil {
				http.Error(w, "Appointment already exists", http.StatusConflict)
				return
			}
			appt.CreateUuid()
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
			author_id := r.URL.Query().Get("author_id")
			if author_id == "" {
				http.Error(w, "Author parameter required", http.StatusBadRequest)
				return
			}
			start_time := r.URL.Query().Get("start_time")
			if start_time == "" {
				http.Error(w, "StarTime parameter required", http.StatusBadRequest)
				return
			}
			var appt api.Appointment
			row1 := db.QueryRow("SELECT * FROM appointments WHERE author_id = $1 AND start_time = $2", author_id, start_time)
			err1 := row1.Scan(&appt.Id, &appt.Title, &appt.Location, &appt.Description, &appt.StartTime, &appt.EndTime, &appt.AuthorId)
			if err1 != nil {
				if err1 == sql.ErrNoRows {
					http.Error(w, "No appointment found", http.StatusNotFound)
				} else {
					http.Error(w, err1.Error(), http.StatusInternalServerError)
				}
				return
			}
			// Fetch participants for the appointment
			var parts_id []string
			rows, err2 := db.Query("SELECT u.id, u.name, u.email FROM users u JOIN appointment_participants ap ON u.id = ap.user_id	WHERE ap.appointment_id = $1", appt.Id)
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()
			for rows.Next() {
				var participant api.User
				if err := rows.Scan(&participant.Id, &participant.Name, &participant.Email); err != nil {
					http.Error(w, "Error reading participant data", http.StatusInternalServerError)
					return
				}
				parts_id = append(parts_id, participant.Id)
			}
			if err := rows.Err(); err != nil {
				http.Error(w, "Error iterating participant data", http.StatusInternalServerError)
				return
			}
			appt.Parts_id = parts_id
			json.NewEncoder(w).Encode(appt)

		case http.MethodPut: // Update
		}
	}
}

// type User struct {
// 	Id       string `json:"id"`
// 	Name     string `json:"name"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type Appointment struct {
// 	Id          string   `json:"id"`
// 	Title       string   `json:"title"`
// 	Location    string   `json:"location"`
// 	Description string   `json:"description"`
// 	StartTime   int      `json:"start_time"`
// 	EndTime     int      `json:"end_time"`
// 	AuthorId    string   `json:"author_id"`
// 	Parts_id    []string `json:"parts_id"`
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
