package main

import (
	"database/sql"

	adapterDb "github.com/GabrielBrotas/hexagonal-architeture/adapters/db"
	"github.com/GabrielBrotas/hexagonal-architeture/application"
	_ "github.com/mattn/go-sqlite3"
)


func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")

	productDbAdapter := adapterDb.NewProductDb(db)

	productService := application.NewProductService(productDbAdapter)

	product, _ := productService.Create("abc", 30)

	product.Enable()

	productService.Enable(product)
}