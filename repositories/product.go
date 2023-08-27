package repositories

import (
	"context"
	"database/sql"

	"github.com/mochammadshenna/ruwar-app/entity"
)

type repository struct {
	db *sql.DB
}

func (repository *repository) GetAllProducts(ctx context.Context, tx *sql.Tx) ([]*entity.Product, error) {
	var products Product


	res := []*entity.Product{}

	res = append(res,products.Entity())


	return res,nil
}