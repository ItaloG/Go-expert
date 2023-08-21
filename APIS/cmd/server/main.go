package main

import (
	"net/http"

	"github.com/ItaloG/Go-expert/APIS/configs"
	"github.com/ItaloG/Go-expert/APIS/internal/entity"
	"github.com/ItaloG/Go-expert/APIS/internal/infra/database"
	"github.com/ItaloG/Go-expert/APIS/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProducts)
	r.Put("/products/{id}", productHandler.UpdateProduct)

	http.ListenAndServe(":8000", r)
}
