package main

import (
	"github.com/ActionKlo/test-ejaw/config"
	"github.com/ActionKlo/test-ejaw/internal/middleware"
	"github.com/ActionKlo/test-ejaw/logger"
	"net/http"
)

func main() {
	log := logger.InitLogger()
	cfg := config.InitConfig(log)

	services := cfg.InitServices(log)
	h := services.H

	mux := http.NewServeMux()

	mux.HandleFunc("GET /sellers", middleware.Auth(h.GetSellers))
	mux.HandleFunc("POST /seller", middleware.Auth(h.InsertSeller))
	mux.HandleFunc("DELETE /seller/{id}", middleware.Auth(h.DeleteSeller))
	mux.HandleFunc("PUT /seller/{id}", middleware.Auth(h.UpdateSeller))

	log.Info("Server started on port: 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Failed to start server")
	}
}
