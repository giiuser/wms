// model.go

package model

type ReceiptRow struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ReceiptId int    `json:"receipt_id"`
	ProductId int    `json:"product_id"`
	Qty       int    `json:"qty"`
}

func CreateReceiptRow(receiptId int, productId int, qty int) (ReceiptRow, error) {
	var rr ReceiptRow

	err := DB.QueryRow(
		"INSERT INTO receipt_table(receipt_id, product_id, qty) VALUES($1, $2, $3) RETURNING id, receipt_id, product_id, qty",
		receiptId, productId, qty).Scan(&rr.ID, &rr.ReceiptId, &rr.ProductId, &rr.Qty)

	if err != nil {
		return rr, err
	}

	return rr, nil
}

func DeleteReceiptRow(id int) error {
	_, err := DB.Exec("DELETE FROM receipt_table WHERE id=$1", id)

	return err
}
