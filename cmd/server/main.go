package main

import (
	"net/http"

	"github.com/thenopholo/back_commerce/configs"

	"github.com/thenopholo/back_commerce/internal/entity"
	"github.com/thenopholo/back_commerce/internal/infra/database"
	"github.com/thenopholo/back_commerce/internal/infra/webserver/handler"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
