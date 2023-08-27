package services

import (
	"context"
	"database/sql"

	"github.com/mochammadshenna/ruwar-app/services/api"
)

type service struct{
	adapter Adapter
	DB *sql.DB
}

func NewService(adapter Adapter, DB *sql.DB) Service {
	return &service{
        adapter: adapter,
		DB: DB,
    }
}

func (service *service)GetAllProducts(ctx context.Context)(*api.FindAllProductResponse, error){
	var res = &api.FindAllProductResponse{}
	return res, nil
}