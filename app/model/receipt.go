// model.go

package model

// var DB *sql.DB

type Receipt struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type ReceiptRow struct {
	ReceiptId int `json:"receipt_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

func GetReceipt(id int) (Receipt, error) {
	var r Receipt
	row := DB.QueryRow("SELECT id, name FROM receipt WHERE id=$1", id)

	if err := row.Scan(&r.ID, &r.Status); err != nil {
		return r, err
	}

	return r, nil
}

func UpdateReceipt(id int, name string) error {
	_, err := DB.Exec("UPDATE receipt SET name=$1 WHERE id=$2",
		name, id)

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

func CreateReceiptRow(receiptId int, productId int, qty int) (ReceiptRow, error) {
	var rr ReceiptRow

	err := DB.QueryRow(
		"INSERT INTO receipt_table(receipt_id, product_id, qty) VALUES($1, $2, $3) RETURNING receipt_id, product_id, qty",
		receiptId, productId, qty).Scan(&rr.ReceiptId, &rr.ProductId, &rr.Qty)

	if err != nil {
		return rr, err
	}

	return rr, nil
}

func GetReceipts(start, count int) ([]Receipt, error) {
	rows, err := DB.Query(
		"SELECT id, name FROM receipt LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Receipt{}

	for rows.Next() {
		var p Receipt
		if err := rows.Scan(&p.ID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func ChangeStatusReceipt(id int, status int) error {
	_, err := DB.Exec("UPDATE receipt SET status=$1 WHERE id=$2",
		status, id)

	return err
}
