package handlers

import (
	"encoding/json"
	"errors"
	"github.com/ActionKlo/test-ejaw/internal/data"
	"github.com/ActionKlo/test-ejaw/internal/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
)

func GetSellers(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*zap.Logger)

	sellers, err := data.Seller{}.GetAll()
	if err != nil {
		logger.Error("failed to get sellers", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("got all sellers", zap.Any("sellers", sellers))
	utils.SendJSON(w, http.StatusOK, sellers)
}

func InsertSeller(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*zap.Logger)

	var seller data.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		logger.Error("failed to Unmarshal seller data", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("failed to close Body", zap.Error(err))
		}
	}(r.Body)

	if seller.Name == "" || seller.Phone == "" {
		logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	id, err := seller.Insert()
	if err != nil {
		logger.Error("failed to insert seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	logger.Info("seller inserted", zap.Any("seller ID", id))
	utils.SendJSON(w, http.StatusCreated, struct {
		ID int `json:"id"`
	}{
		ID: id,
	})
}

func DeleteSeller(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*zap.Logger)

	var seller data.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		logger.Error("failed to Unmarshal seller data", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("failed to close Body", zap.Error(err))
		}
	}(r.Body)

	if seller.ID <= 0 {
		logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	deleted, err := seller.Delete()
	if err != nil {
		logger.Error("failed to delete seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !deleted {
		logger.Error("failed to delete seller", zap.Error(errors.New("seller not found")))
		http.Error(w, errors.New("seller not found").Error(), http.StatusBadRequest)
	}

	logger.Info("seller deleted", zap.Any("seller ID", seller.ID))
	utils.SendJSON(w, http.StatusOK, struct {
		ID int `json:"id"`
	}{
		ID: seller.ID,
	})
}

func UpdateSeller(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value("logger").(*zap.Logger)

	var seller data.Seller
	if err := json.NewDecoder(r.Body).Decode(&seller); err != nil {
		logger.Error("failed to Unmarshal seller data", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Error("failed to close Body", zap.Error(err))
		}
	}(r.Body)

	if seller.ID <= 0 || seller.Name == "" || seller.Phone == "" {
		logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	updated, err := seller.Update()
	if err != nil {
		logger.Error("failed to update seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !updated {
		logger.Error("failed to update seller", zap.Error(errors.New("seller not found")))
		http.Error(w, errors.New("seller not fount").Error(), http.StatusBadRequest)
	}

	logger.Info("seller updated", zap.Any("seller ID", seller.ID))
	utils.SendJSON(w, http.StatusOK, struct {
		ID int `json:"id"`
	}{
		ID: seller.ID,
	})
}
