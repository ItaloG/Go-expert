package main

import (
	"database/sql"
	"fmt"

	"github.com/ItaloG/Go-expert/DI/product"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	// Create a new product repository
	// repository := product.NewProductRepository(db)

	// // Create a new product usecase
	// usecase := product.NewProductUseCase(repository)

	// with wire
	usecase := NewUseCase(db)

	product, err := usecase.GetProduct(1)
	if err != nil {
		panic(err)
	}

	fmt.Println(product.Name)
}
