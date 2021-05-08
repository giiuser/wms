package model

import (
	"database/sql"
	"time"
)

type Allocation struct {
	BaseModel
	Status       int             `json:"status"`
	DocumentId   int             `json:"document_id"`
	DocumentType string          `json:"document_type"`
	Products     []AllocationRow `json:"products"`
}

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

func GetAllocation(id int) (Allocation, error) {
	var a Allocation
	row := DB.QueryRow("SELECT id, status, document_id, document_type FROM allocation WHERE id=$1", id)

	if err := row.Scan(&a.ID, &a.Status, &a.DocumentId, &a.DocumentType); err != nil {
		return a, err
	}

	rows, err := DB.Query(
		"SELECT at.id, p.name, p.brand, at.allocation_id, at.product_id, at.qty, at.cell_id, c.name FROM allocation_table at JOIN product p ON p.id = at.product_id LEFT JOIN cell c ON c.id = at.cell_id WHERE allocation_id=$1",
		id)

	if err != nil {
		return a, err
	}

	defer rows.Close()

	products := []AllocationRow{}

	for rows.Next() {
		var ar AllocationRow
		if err := rows.Scan(&ar.ID, &ar.Name, &ar.Brand, &ar.AllocationId, &ar.ProductId, &ar.Qty, &ar.CellId, &ar.CellName); err != nil {
			return a, err
		}
		products = append(products, ar)
	}
	a.Products = products

	return a, nil
}

func GetAllocations(start, count int) ([]Allocation, error) {
	rows, err := DB.Query(
		"SELECT id, status, document_id, document_type, created_at, updated_at FROM allocation LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allocations := []Allocation{}

	for rows.Next() {
		var a Allocation
		if err := rows.Scan(&a.ID, &a.Status, &a.DocumentId, &a.DocumentType, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		allocations = append(allocations, a)
	}

	return allocations, nil
}

func UpdateAllocation(id int, documentId int, documentType string) error {
	_, err := DB.Exec("UPDATE allocation SET document_id=$1, document_type=$2, updated_at=$3 WHERE id=$4",
		documentId, documentType, time.Now(), id)

	return err
}

func DeleteAllocation(id int) error {
	_, err := DB.Exec("DELETE FROM allocation WHERE id=$1", id)

	return err
}

func CreateAllocation(documentId int, documentType string) (Allocation, error) {
	var a Allocation

	err := DB.QueryRow(
		"INSERT INTO allocation(document_id, document_type) VALUES($1, $2) RETURNING id, status, document_id, document_type",
		documentId, documentType).Scan(&a.ID, &a.Status, &a.DocumentId, &a.DocumentType)

	if err != nil {
		return a, err
	}

	return a, nil
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

func ChangeStatusAllocation(id int, status int) error {
	_, err := DB.Exec("UPDATE allocation SET status=$1, updated_at=$2 WHERE id=$3",
		status, time.Now(), id)

	if status == 2 {
		MakeCellPosting(id, "allocation", true)
	} else if status == 1 {
		MakeCellPosting(id, "allocation", false)
	}

	return err
}

func DeleteAllocationRows(id int) error {
	_, err := DB.Exec("DELETE FROM allocation_table WHERE allocation_id=$1", id)

	return err
}
