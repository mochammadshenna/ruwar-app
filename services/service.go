package services

import (
	"context"

	"github.com/mochammadshenna/ruwar-app/services/api"
)

type Service interface {
	GetAllProducts(ctx context.Context)(*api.FindAllProductResponse, error) 
}