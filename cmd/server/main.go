package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	UserDB := database.NewUser(db)
	userHandler := handler.NewUserHandler(UserDB)

	r:= chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)
	r.Get("/products", productHandler.GetProducts)

	r.Post("/users", userHandler.CreateUser)

	http.ListenAndServe(":8000", r)
}
