package repository

import (
	"github.com/anvi23mth/inventory-system/internal/database"
	"github.com/anvi23mth/inventory-system/internal/model"
)

func CreateProduct(p model.Product) error {

	query := `
	INSERT INTO products(id,name,description,category,price,brand,quantity)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`

	_, err := database.DB.Exec(
		query,
		p.ID,
		p.Name,
		p.Description,
		p.Category,
		p.Price,
		p.Brand,
		p.Quantity,
	)

	return err
}

func GetProducts() ([]model.Product, error) {

	rows, err := database.DB.Query(`SELECT * FROM products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {

		var p model.Product

		rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Category,
			&p.Price,
			&p.Brand,
			&p.Quantity,
		)

		products = append(products, p)
	}

	return products, nil
}
func GetProductByID(id string) (model.Product, error) {

	query := `SELECT * FROM products WHERE id=$1`

	row := database.DB.QueryRow(query, id)

	var p model.Product

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Category,
		&p.Price,
		&p.Brand,
		&p.Quantity,
	)

	return p, err
}
func UpdateProduct(id string, p model.Product) error {

	query := `
	UPDATE products
	SET name=$1, description=$2, category=$3, price=$4, brand=$5, quantity=$6
	WHERE id=$7
	`

	_, err := database.DB.Exec(
		query,
		p.Name,
		p.Description,
		p.Category,
		p.Price,
		p.Brand,
		p.Quantity,
		id,
	)

	return err
}
func DeleteProduct(id string) error {

	query := `DELETE FROM products WHERE id=$1`

	_, err := database.DB.Exec(query, id)

	return err
}
