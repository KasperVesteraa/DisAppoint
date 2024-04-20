package server

import (
	"database/sql"
	"net/http"
)

func InitializeRoutes(db *sql.DB) {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/users", UserHandler(db))
	http.HandleFunc("/appointments", AppointmentHandler(db))
}
