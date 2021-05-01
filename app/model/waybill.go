package model

type Waybill struct {
	ID           int    `json:"id"`
	Status       int    `json:"status"`
	DocumentId   int    `json:"document_id"`
	DocumentType string `json:"document_type"`
}

type WaybillRow struct {
	WaybillId int `json:"waybill_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

func GetWaybill(id int) (Receipt, error) {
	var r Receipt
	row := DB.QueryRow("SELECT id, name FROM waybill WHERE id=$1", id)

	if err := row.Scan(&r.ID, &r.Status); err != nil {
		return r, err
	}

	return r, nil
}

func UpdateWaybill(id int, name string) error {
	_, err := DB.Exec("UPDATE waybill SET name=$1 WHERE id=$2",
		name, id)

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

func GetWaybills(start, count int) ([]Receipt, error) {
	rows, err := DB.Query(
		"SELECT id, name FROM waybill LIMIT $1 OFFSET $2",
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

func ChangeStatusWaybill(id int, status int) error {
	_, err := DB.Exec("UPDATE waybill SET status=$1 WHERE id=$2",
		status, id)

	if status == 2 {
		MakePosting(id, "waybill", false)
	} else if status == 1 {
		MakePosting(id, "waybill", true)
	}

	return err
}
