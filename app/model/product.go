// model.go

package model

import (
	"database/sql"
)

var DB *sql.DB

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
}

func GetProduct(id int) (Product, error) {
	var p Product
	row := DB.QueryRow("SELECT id, name, brand FROM product WHERE id=$1", id)

	if err := row.Scan(&p.ID, &p.Name, &p.Brand); err != nil {
		return p, err
	}

	return p, nil
}

func UpdateProduct(id int, name string, brand string) error {
	_, err := DB.Exec("UPDATE product SET name=$1, brand=$2 WHERE id=$3",
		name, brand, id)

	return err
}

func DeleteProduct(id int) error {
	_, err := DB.Exec("DELETE FROM product WHERE id=$1", id)

	return err
}

func CreateProduct(name string, brand string) (Product, error) {
	var p Product
	err := DB.QueryRow(
		"INSERT INTO product(name, brand) VALUES($1, $2) RETURNING id, name, brand",
		name, brand).Scan(&p.ID, &p.Name, &p.Brand)

	if err != nil {
		return p, err
	}

	return p, nil
}

func GetProducts(start, count int) ([]Product, error) {
	rows, err := DB.Query(
		"SELECT id, name, brand FROM product LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Brand); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
