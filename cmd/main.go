package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	config "github.com/mochammadshenna/ruwar-app/configs"
	"github.com/mochammadshenna/ruwar-app/controllers"
	db "github.com/mochammadshenna/ruwar-app/db"
	"github.com/mochammadshenna/ruwar-app/routers"
	"github.com/mochammadshenna/ruwar-app/services"
)

func main (){
	// validate := validator.New()
	config.Init()
	db := db.NewDB()

	productRepository := services.Adapter{}
	productService := services.NewService(productRepository, db)
	productController := controllers.NewProductController(productService)
	router := routers.NewRouter(productController)

	host := fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port)
	server := &http.Server{
		Addr:    host,
		Handler: router,
	}

	fmt.Printf("Apps running on %s \n", host)
	err := server.ListenAndServe()
	if err!= nil {
      panic(err)
    }
}