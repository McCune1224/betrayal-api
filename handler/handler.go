package handler

import (
	"github.com/jmoiron/sqlx"
	"github.com/mccune1224/betrayal-api/models"
)

type Handler struct {
	models *models.Models
}

func NewHandler(db *sqlx.DB) *Handler {
	model := models.NewModels(db)
	return &Handler{
		models: model,
	}
}
