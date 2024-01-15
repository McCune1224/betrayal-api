package models

import (
	"github.com/jmoiron/sqlx"
)

type Models struct {
	db        *sqlx.DB
	ItemModel *ItemModel
}

func NewModels(Db *sqlx.DB) *Models {
	return &Models{
		db:        Db,
		ItemModel: &ItemModel{Db: Db},
	}
}
