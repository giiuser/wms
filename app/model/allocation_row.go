package model

import (
	"database/sql"
)

type AllocationRow struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`
	Brand        string         `json:"brand"`
	AllocationId int            `json:"allocation_id"`
	ProductId    int            `json:"product_id"`
	Qty          int            `json:"qty"`
	CellId       int            `json:"cell_id"`
	CellName     sql.NullString `json:"cell_name"`
}

func CreateAllocationRow(waybillId int, productId int, qty int, cellId int) (AllocationRow, error) {
	var ar AllocationRow

	err := DB.QueryRow(
		"INSERT INTO allocation_table(allocation_id, product_id, qty, cell_id) VALUES($1, $2, $3, $4) RETURNING id, allocation_id, product_id, qty, cell_id",
		waybillId, productId, qty, cellId).Scan(&ar.ID, &ar.AllocationId, &ar.ProductId, &ar.Qty, &ar.CellId)

	if err != nil {
		return ar, err
	}

	return ar, nil
}

func UpdateAllocationRow(id int, cellId int) error {
	_, err := DB.Exec("UPDATE allocation_table SET cell_id=$1 WHERE id=$2",
		cellId, id)

	return err
}

func DeleteAllocationRows(id int) error {
	_, err := DB.Exec("DELETE FROM allocation_table WHERE allocation_id=$1", id)

	return err
}
