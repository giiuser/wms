// model.go

package model

import (
	"time"
)

type Receipt struct {
	BaseModel
	Status   int          `json:"status"`
	Products []ReceiptRow `json:"products"`
}

func GetReceipt(id int) (Receipt, error) {
	var r Receipt
	row := DB.QueryRow("SELECT id, status FROM receipt WHERE id=$1", id)

	if err := row.Scan(&r.ID, &r.Status); err != nil {
		return r, err
	}

	rows, err := DB.Query(
		"SELECT rt.id, p.name, rt.receipt_id, rt.product_id, rt.qty FROM receipt_table rt JOIN product p ON p.id = rt.product_id WHERE receipt_id=$1",
		id)

	if err != nil {
		return r, err
	}

	defer rows.Close()

	products := []ReceiptRow{}

	for rows.Next() {
		var rr ReceiptRow
		if err := rows.Scan(&rr.ID, &rr.Name, &rr.ReceiptId, &rr.ProductId, &rr.Qty); err != nil {
			return r, err
		}
		products = append(products, rr)
	}
	r.Products = products

	return r, nil
}

func GetReceipts(start, count int) ([]Receipt, error) {
	rows, err := DB.Query(
		"SELECT id, status, created_at, updated_at FROM receipt LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	receipts := []Receipt{}

	for rows.Next() {
		var r Receipt
		if err := rows.Scan(&r.ID, &r.Status, &r.CreatedAt, &r.UpdatedAt); err != nil {
			return nil, err
		}
		receipts = append(receipts, r)
	}

	return receipts, nil
}

func UpdateReceipt(id int, name string) error {
	_, err := DB.Exec("UPDATE receipt SET name=$1, updated_at=$2 WHERE id=$2",
		name, time.Now(), id)

	return err
}

func DeleteReceipt(id int) error {
	_, err := DB.Exec("DELETE FROM receipt WHERE id=$1", id)

	return err
}

func CreateReceipt() (Receipt, error) {
	var r Receipt

	err := DB.QueryRow(
		"INSERT INTO receipt DEFAULT VALUES RETURNING id, status").Scan(&r.ID, &r.Status)

	if err != nil {
		return r, err
	}

	return r, nil
}

func ChangeStatusReceipt(id int, status int) error {
	_, err := DB.Exec("UPDATE receipt SET status=$1, updated_at=$2 WHERE id=$3",
		status, time.Now(), id)

	if status == 2 {
		MakePosting(id, "receipt", true)
	} else if status == 1 {
		MakePosting(id, "receipt", false)
	}

	return err
}
