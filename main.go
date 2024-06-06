package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"hello/internal/api/handlers"
	"hello/internal/api/middleware"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello")

	basicAuth := middleware.BasicAuth()
	router := mux.NewRouter()

	router.Handle("/api/register", handlers.Register()).Methods("POST")
	router.Handle("/api/login", handlers.Login()).Methods("POST")

	router.Handle("/api/posts", basicAuth(handlers.CreatePost())).Methods("POST")
	router.Handle("/api/posts", basicAuth(handlers.ListOfPosts())).Methods("GET")
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
