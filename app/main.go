package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to Ecommerce DevOps Platform")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Application Healthy")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}