package model

import (
	"time"
)

type Waybill struct {
	BaseModel
	Status       int          `json:"status"`
	DocumentId   int          `json:"document_id"`
	DocumentType string       `json:"document_type"`
	Products     []WaybillRow `json:"products"`
}

type WaybillRow struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	WaybillId int    `json:"waybill_id"`
	ProductId int    `json:"product_id"`
	Qty       int    `json:"qty"`
}

func GetWaybill(id int) (Waybill, error) {
	var wb Waybill
	row := DB.QueryRow("SELECT id, status, document_id, document_type FROM waybill WHERE id=$1", id)

	if err := row.Scan(&wb.ID, &wb.Status, &wb.DocumentId, &wb.DocumentType); err != nil {
		return wb, err
	}

	rows, err := DB.Query(
		"SELECT wt.id, p.name, p.brand, wt.waybill_id, wt.product_id, wt.qty FROM waybill_table wt JOIN product p ON p.id = wt.product_id WHERE waybill_id=$1",
		id)

	if err != nil {
		return wb, err
	}

	defer rows.Close()

	products := []WaybillRow{}

	for rows.Next() {
		var wbr WaybillRow
		if err := rows.Scan(&wbr.ID, &wbr.Name, &wbr.Brand, &wbr.WaybillId, &wbr.ProductId, &wbr.Qty); err != nil {
			return wb, err
		}
		products = append(products, wbr)
	}
	wb.Products = products

	return wb, nil
}

func UpdateWaybill(id int, documentId int, documentType string) error {
	_, err := DB.Exec("UPDATE waybill SET document_id=$1, document_type=$2, updated_at=$3 WHERE id=$4",
		documentId, documentType, time.Now(), id)

	return err
}

func DeleteWaybill(id int) error {
	_, err := DB.Exec("DELETE FROM waybill WHERE id=$1", id)

	return err
}

func CreateWaybill(documentId int, documentType string) (Waybill, error) {
	var w Waybill

	err := DB.QueryRow(
		"INSERT INTO waybill(document_id, document_type) VALUES($1, $2) RETURNING id, status, document_id, document_type",
		documentId, documentType).Scan(&w.ID, &w.Status, &w.DocumentId, &w.DocumentType)

	if err != nil {
		return w, err
	}

	return w, nil
}

func CreateWaybillRow(waybillId int, productId int, qty int) (WaybillRow, error) {
	var wbr WaybillRow

	err := DB.QueryRow(
		"INSERT INTO waybill_table(waybill_id, product_id, qty) VALUES($1, $2, $3) RETURNING waybill_id, product_id, qty",
		waybillId, productId, qty).Scan(&wbr.WaybillId, &wbr.ProductId, &wbr.Qty)

	if err != nil {
		return wbr, err
	}

	return wbr, nil
}

func GetWaybills(start, count int) ([]Waybill, error) {
	rows, err := DB.Query(
		"SELECT id, status, document_id, document_type, created_at, updated_at FROM waybill LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Waybill{}

	for rows.Next() {
		var wb Waybill
		if err := rows.Scan(&wb.ID, &wb.Status, &wb.DocumentId, &wb.DocumentType, &wb.CreatedAt, &wb.UpdatedAt); err != nil {
			return nil, err
		}
		products = append(products, wb)
	}

	return products, nil
}

func ChangeStatusWaybill(id int, status int) error {
	_, err := DB.Exec("UPDATE waybill SET status=$1, updated_at=$2 WHERE id=$3",
		status, time.Now(), id)

	if status == 2 {
		MakePosting(id, "waybill", false)
	} else if status == 1 {
		MakePosting(id, "waybill", true)
	}

	return err
}

func DeleteWaybillRows(id int) error {
	_, err := DB.Exec("DELETE FROM waybill_table WHERE waybill_id=$1", id)

	return err
}
