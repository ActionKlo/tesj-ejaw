package handlers

import (
	"github.com/ActionKlo/test-ejaw/internal/repository"
	"go.uber.org/zap"
)

type RecipesHandler struct {
	logger *zap.Logger
	store  *repository.Service
}

func InitRecipesHandler(logger *zap.Logger, store *repository.Service) *RecipesHandler {
	return &RecipesHandler{
		logger: logger,
		store:  store,
	}
}
