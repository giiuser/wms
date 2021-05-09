// model.go

package model

import (
	"database/sql"
)

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

func DeleteCollectRow(id int) error {
	_, err := DB.Exec("DELETE FROM collect_table WHERE id=$1", id)

	return err
}
