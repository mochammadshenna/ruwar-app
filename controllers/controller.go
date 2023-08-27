package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ProductControllers interface{
	FindAllProducts(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}