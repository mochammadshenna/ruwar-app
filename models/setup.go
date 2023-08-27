package models

import (
	"log"

	"github.com/mochammadshenna/ruwar-app/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	database, err := gorm.Open(postgres.Open("postgres:postgres@tcp(localhost:5432)/ruwar_db"))
	if err != nil {
		log.Fatal("Failed to open database")
		panic(err)
	}

	database.AutoMigrate(&entity.Product{})

	DB = database;
}