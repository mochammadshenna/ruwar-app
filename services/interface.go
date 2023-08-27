package services

import (
	"context"
	"database/sql"
)

type Adapter struct {
	ProductRepository ProductRepository
}

type ProductRepository interface {
	GetAllProducts(ctx context.Context, tx *sql.Tx)
}