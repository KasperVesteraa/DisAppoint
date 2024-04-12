package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/KasperVesteraa/DisAppoint/internal/server"
	_ "github.com/lib/pq"
)

const (
	Host     = "localhost"
	Port     = 5434
	User     = "root"
	Password = "root"
	Dbname   = "disappoint_db"
)

var db *sql.DB

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database.")

	server.InitializeRoutes(db)
	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
