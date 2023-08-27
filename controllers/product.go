package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/ruwar-app/services"
)

type ProductController struct {
	ProductService services.Service
}

func NewProductController(productService services.Service) ProductControllers{
	return &ProductController{
		ProductService: productService,
	}
}

func (controller *ProductController)FindAllProducts(writer http.ResponseWriter, request *http.Request, param httprouter.Params){
	response , err := controller.ProductService.GetAllProducts(request.Context())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(response)
	// httphelper.Write(request.Context(), writer, response)
	writer.Write([]byte("Done"))
}