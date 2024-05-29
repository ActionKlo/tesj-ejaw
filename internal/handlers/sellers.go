package handlers

import (
	"encoding/json"
	"github.com/ActionKlo/test-ejaw/internal/data"
	"net/http"
)

func GetSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := data.Seller{}.GetAll()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sellers)
}
