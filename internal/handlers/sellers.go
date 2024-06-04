package handlers

import (
	"encoding/json"
	"errors"
	"github.com/ActionKlo/test-ejaw/internal/models"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

func sendJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")

	if statusCode != http.StatusOK {
		w.WriteHeader(statusCode)
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *RecipesHandler) GetSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := h.store.GetAll()
	if err != nil {
		h.logger.Error("failed to get sellers", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("got all sellers", zap.Any("sellers", sellers))
	sendJSON(w, http.StatusOK, sellers)
}

func (h *RecipesHandler) InsertSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("failed to read seller repository", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			h.logger.Error("failed to close Body", zap.Error(err))
		}
	}(r.Body)

	if err = json.Unmarshal(body, &seller); err != nil {
		h.logger.Error("failed to Unmarshal seller repository", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if seller.Name == "" || seller.Phone == "" {
		h.logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	id, err := h.store.Insert(seller.Name, seller.Phone)
	if err != nil {
		h.logger.Error("failed to insert seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("seller inserted", zap.Any("seller ID", id))
	sendJSON(w, http.StatusCreated, struct {
		ID int `json:"id"`
	}{
		ID: id,
	})
}

func (h *RecipesHandler) DeleteSeller(w http.ResponseWriter, r *http.Request) {
	sellerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || sellerID <= 0 {
		h.logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	deleted, err := h.store.Delete(sellerID)
	if err != nil {
		h.logger.Error("failed to delete seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !deleted {
		h.logger.Error("failed to delete seller", zap.Error(errors.New("seller not found")))
		http.Error(w, errors.New("seller not found").Error(), http.StatusBadRequest)
	}

	h.logger.Info("seller deleted", zap.Any("seller ID", sellerID))
	sendJSON(w, http.StatusOK, struct {
		ID int `json:"id"`
	}{
		ID: sellerID,
	})
}

func (h *RecipesHandler) UpdateSeller(w http.ResponseWriter, r *http.Request) {
	var seller models.Seller
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("failed to Unmarshal seller repository", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			h.logger.Error("failed to close Body", zap.Error(err))
		}
	}(r.Body)

	if err = json.Unmarshal(body, &seller); err != nil {
		h.logger.Error("failed to Unmarshal seller repository", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sellerID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || sellerID <= 0 {
		h.logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	if sellerID <= 0 || seller.Name == "" || seller.Phone == "" {
		h.logger.Error("bad request body", zap.Error(errors.New("BadRequest")))
		http.Error(w, errors.New("bad request body").Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.store.Update(sellerID, seller.Name, seller.Phone)
	if err != nil {
		h.logger.Error("failed to update seller", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !updated {
		h.logger.Error("failed to update seller", zap.Error(errors.New("seller not found")))
		http.Error(w, errors.New("seller not fount").Error(), http.StatusBadRequest)
	}

	h.logger.Info("seller updated", zap.Any("seller ID", seller.ID))
	sendJSON(w, http.StatusOK, struct {
		ID int `json:"id"`
	}{
		ID: seller.ID,
	})
}
