package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to DisAppoint!\n")
	})

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
	
