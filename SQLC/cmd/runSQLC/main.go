package main

import (
	"context"
	"database/sql"

	"github.com/ItaloG/Go-expert/SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Backend",
	// 	Description: sql.NullString{String: "Backend description", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)
	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.ID, category.Name, category.Description.String)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "a145c0f8-2563-4f88-9fd5-3e0abdfd4e77",
	// 	Name:        "Backend updated",
	// 	Description: sql.NullString{String: "Backend description updated", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		println(category.ID, category.Name, category.Description.String)
	}

	err = queries.DeleteCategory(ctx, "a145c0f8-2563-4f88-9fd5-3e0abdfd4e77")
	if err != nil {
		panic(err)
	}
}
