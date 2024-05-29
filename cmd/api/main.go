package main

import (
	"github.com/ActionKlo/test-ejaw/internal/data"
	"github.com/ActionKlo/test-ejaw/internal/handlers"
	"log"
	"net/http"
)

func main() {
	data.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /sellers", handlers.GetSellers)

	log.Println("Server started on port: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server")
	}
}
