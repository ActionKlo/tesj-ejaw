package handlers

import (
	"encoding/json"
	"errors"
	"github.com/ActionKlo/test-ejaw/internal/data"
	"github.com/ActionKlo/test-ejaw/internal/utils"
	"io"
	"log"
	"net/http"
)

func GetSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := data.Seller{}.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendJSON(w, sellers)
}

func InsertSeller(w http.ResponseWriter, r *http.Request) {
	var seller data.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("failed to close r.Body:", err.Error())
		}
	}(r.Body)

	if seller.Name == "" || seller.Phone == "" {
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	id, err := seller.Insert()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendJSON(w, struct {
		ID int `json:"id"`
	}{
		ID: id,
	})
}
