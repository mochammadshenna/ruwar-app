package api

type FindAllProductResponse struct {
	List []ProductResponse `json:"list"`
}

type ProductResponse struct {
	Id int64 `json:"id"`
	ProductName string `json:"product_name"`
	Description string `json:"description"`
}