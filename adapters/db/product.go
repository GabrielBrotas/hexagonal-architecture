package db

import (
	"database/sql"
	"github.com/GabrielBrotas/hexagonal-architeture/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) GetById(id string) (application.IProduct, error) {
	var product application.Product

	statement, err := p.db.Prepare("select id, name, price, status from products where id=?")

	if err != nil {
		return nil, err
	}

	err = statement.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product application.IProduct) (application.IProduct, error) {
	var rows int
	
	// p.db.QueryRow("select id from products where id = ?", product.GetID()).Scan(&rows)

	err := p.db.QueryRow("select count(*) from products where id = ?", product.GetID()).Scan(&rows)

	if err != nil {
		return nil, err
	}

	if rows == 0 {
		_, err := p.create(product)

		if err != nil {
			return nil, err
		}

		return product, nil
	} else {
		_, err := p.update(product)
	
		if err != nil {
			return nil, err
		}
			
		return product, nil
	}

}

// metodo privado
func (p *ProductDb) create(product application.IProduct) (application.IProduct, error) {
	statement, err := p.db.Prepare("insert into products(id, name, price, status) values(?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = statement.Exec(
		product.GetID(), 
		product.GetName(), 
		product.GetPrice(), 
		product.GetStatus(),
	)

	if err != nil {
		return nil, err
	}

	err = statement.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product application.IProduct) (application.IProduct, error) {
	_, err := p.db.Exec("update products set name = ?, price = ?, status = ? where id = ?",
		product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
