// model.go

package model

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetProduct(id int) (Product, error) {
	var p Product
	row := DB.QueryRow("SELECT id, name FROM product WHERE id=$1", id)

	if err := row.Scan(&p.ID, &p.Name); err != nil {
		return p, err
	}

	return p, nil
}

func UpdateProduct(id int, name string) error {
	_, err := DB.Exec("UPDATE product SET name=$1 WHERE id=$2",
		name, id)

	return err
}

func DeleteProduct(id int) error {
	_, err := DB.Exec("DELETE FROM product WHERE id=$1", id)

	return err
}

func CreateProduct(name string) (Product, error) {
	var p Product
	fmt.Println(name)
	err := DB.QueryRow(
		"INSERT INTO product(name) VALUES($1) RETURNING id, name",
		name).Scan(&p.ID, &p.Name)

	if err != nil {
		return p, err
	}

	return p, nil
}

func GetProducts(start, count int) ([]Product, error) {
	rows, err := DB.Query(
		"SELECT id, name FROM product LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
