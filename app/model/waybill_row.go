package model

type WaybillRow struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Brand     string `json:"brand"`
	WaybillId int    `json:"waybill_id"`
	ProductId int    `json:"product_id"`
	Qty       int    `json:"qty"`
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

func DeleteWaybillRows(id int) error {
	_, err := DB.Exec("DELETE FROM waybill_table WHERE waybill_id=$1", id)

	return err
}
