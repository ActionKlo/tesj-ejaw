package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
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

	utils.SendJSON(w, http.StatusOK, sellers)
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

	utils.SendJSON(w, http.StatusCreated, struct {
		ID int `json:"id"`
	}{
		ID: id,
	})
}

func DeleteSeller(w http.ResponseWriter, r *http.Request) {
	// !TODO check if typeOf(id) == int
	var seller data.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(seller)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("failed to close r.Body:", err.Error())
		}
	}(r.Body)

	if seller.ID <= 0 {
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	deleted, err := seller.Delete()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !deleted {
		http.Error(w, errors.New("seller not fount").Error(), http.StatusBadRequest)
	}

	utils.SendJSON(w, http.StatusOK, struct {
		ID int `json:"id"`
	}{
		ID: seller.ID,
	})
}
