package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/GabrielBrotas/hexagonal-architeture/adapters/db"
	"github.com/GabrielBrotas/hexagonal-architeture/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	// banco de dados em memoria
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,	
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `insert into products values("abc", "Product 1", 0, "disabled");`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.GetById("abc")

	require.Nil(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, float64(0), product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product 2"
	product.Price = 25

	resultCreate, err := productDb.Save(product)

	require.Nil(t, err)

	require.Equal(t, product.Name, resultCreate.GetName())
	require.Equal(t, product.Price, resultCreate.GetPrice())
	require.Equal(t, product.Status, resultCreate.GetStatus())

	product.Status = "enabled"
	resultUpdate, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, resultCreate.GetID(), resultUpdate.GetID())
	require.Equal(t, product.Name, resultUpdate.GetName())
	require.Equal(t, product.Price, resultUpdate.GetPrice())
	require.Equal(t, product.Status, resultUpdate.GetStatus())
}

