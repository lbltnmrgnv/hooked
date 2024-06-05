package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"hello/internal/api/handlers"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello")

	router := mux.NewRouter()

	router.HandleFunc("/api/user", handlers.Register).Methods("POST")

	router.Use()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return
	}
}
