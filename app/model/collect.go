// model.go

package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Collect struct {
	BaseModel
	Status   int          `json:"status"`
	Products []CollectRow `json:"products"`
}

type CollectRow struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Brand     string         `json:"brand"`
	CollectId int            `json:"collect_id"`
	ProductId int            `json:"product_id"`
	Qty       int            `json:"qty"`
	CellId    int            `json:"cell_id"`
	CellName  sql.NullString `json:"cell_name"`
}

func GetCollect(id int) (Collect, error) {
	var c Collect
	row := DB.QueryRow("SELECT id, status, created_at FROM collect WHERE id=$1", id)

	if err := row.Scan(&c.ID, &c.Status, &c.CreatedAt); err != nil {
		return c, err
	}

	rows, err := DB.Query(
		"SELECT ct.id, p.name, p.brand, ct.collect_id, ct.product_id, ct.qty, ct.cell_id, c.name FROM collect_table ct JOIN product p ON p.id = ct.product_id LEFT JOIN cell c ON c.id = ct.cell_id WHERE collect_id=$1",
		id)

	if err != nil {
		return c, err
	}

	defer rows.Close()

	products := []CollectRow{}

	for rows.Next() {
		var cr CollectRow
		if err := rows.Scan(&cr.ID, &cr.Name, &cr.Brand, &cr.CollectId, &cr.ProductId, &cr.Qty, &cr.CellId, &cr.CellName); err != nil {
			return c, err
		}
		products = append(products, cr)
	}
	c.Products = products

	return c, nil
}

func GetCollects(start, count int) ([]Collect, error) {
	rows, err := DB.Query(
		"SELECT id, status, created_at, updated_at FROM collect LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	collects := []Collect{}

	for rows.Next() {
		var c Collect
		if err := rows.Scan(&c.ID, &c.Status, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		collects = append(collects, c)
	}

	return collects, nil
}

func UpdateCollect(id int, name string) error {
	_, err := DB.Exec("UPDATE collect SET name=$1, updated_at=$2 WHERE id=$3",
		name, time.Now(), id)

	return err
}

func DeleteCollect(id int) error {
	_, err := DB.Exec("DELETE FROM collect WHERE id=$1", id)

	return err
}

func CreateCollect() (Collect, error) {
	var c Collect

	err := DB.QueryRow(
		"INSERT INTO collect DEFAULT VALUES RETURNING id, status").Scan(&c.ID, &c.Status)

	if err != nil {
		return c, err
	}

	return c, nil
}

func CreateCollectRow(collectId int, productId int, qty int, cell_id int) (CollectRow, error) {
	var cr CollectRow

	err := DB.QueryRow(
		"INSERT INTO collect_table(collect_id, product_id, qty, cell_id) VALUES($1, $2, $3, $4) RETURNING id, collect_id, product_id, qty, cell_id",
		collectId, productId, qty, cell_id).Scan(&cr.ID, &cr.CollectId, &cr.ProductId, &cr.Qty, &cr.CellId)

	if err != nil {
		return cr, err
	}

	return cr, nil
}

func UpdateCollectRow(id int, cellId int) error {
	_, err := DB.Exec("UPDATE collect_table SET cell_id=$1 WHERE id=$2",
		cellId, id)

	return err
}

func ChangeStatusCollect(id int, status int) error {
	fmt.Println("hh")
	_, err := DB.Exec("UPDATE collect SET status=$1, updated_at=$2 WHERE id=$3",
		status, time.Now(), id)

	if status == 2 {
		MakeCellPosting(id, "collect", true)
	} else if status == 1 {
		MakeCellPosting(id, "collect", false)
	}

	return err
}

func DeleteCollectRow(id int) error {
	_, err := DB.Exec("DELETE FROM collect_table WHERE id=$1", id)

	return err
}
