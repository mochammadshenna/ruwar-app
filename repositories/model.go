package repositories

import "github.com/mochammadshenna/ruwar-app/entity"

type Product struct {
	Id int64
	ProductName string
	Description string
}

func (p Product) Entity()*entity.Product{
	return &entity.Product{
        Id: p.Id,
        ProductName: p.ProductName,
        Description: p.Description,
    }
}