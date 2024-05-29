package handlers

import (
	"github.com/ActionKlo/test-ejaw/internal/data"
	"github.com/ActionKlo/test-ejaw/internal/utils"
	"net/http"
)

func GetSellers(w http.ResponseWriter, r *http.Request) {
	sellers, err := data.Seller{}.GetAll()
	if err != nil {
		utils.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.SendJSON(w, sellers)
}
