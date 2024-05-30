package main

import (
	"github.com/ActionKlo/test-ejaw/internal/handlers"
	"github.com/ActionKlo/test-ejaw/internal/middleware"
	"github.com/ActionKlo/test-ejaw/internal/repository"
	"log"
	"net/http"
)

func main() {
	repository.InitDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /sellers", middleware.Logger(middleware.Auth(handlers.GetSellers)))
	mux.HandleFunc("POST /seller", middleware.Logger(middleware.Auth(handlers.InsertSeller)))
	mux.HandleFunc("DELETE /seller/{id}", middleware.Logger(middleware.Auth(handlers.DeleteSeller)))
	mux.HandleFunc("PUT /seller/{id}", middleware.Logger(middleware.Auth(handlers.UpdateSeller)))

	log.Println("Server started on port: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server")
	}
}
