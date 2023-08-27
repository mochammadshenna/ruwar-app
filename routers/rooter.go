package routers

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/ruwar-app/controllers"
)

func NewRouter(productController controllers.ProductControllers) *httprouter.Router{
	router := httprouter.New()

	router.GET("/api/products", productController.FindAllProducts)

	router.PanicHandler = httprouter.New().PanicHandler

	return router
}

