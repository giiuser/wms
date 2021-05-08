// model.go

package model

type Stock struct {
	ProductId int     `json:"status"`
	Qty       int     `json:"qty"`
	Cells     []Cells `json:"cells"`
}

type Cells struct {
	CellName string `json:"cell_name"`
	Qty      int    `json:"qty"`
}

func GetStocks(productId int) (Stock, error) {
	var s Stock
	row := DB.QueryRow("SELECT product_id, SUM(qty) AS qty FROM wh_register WHERE product_id = $1 GROUP BY product_id", productId)

	if err := row.Scan(&s.ProductId, &s.Qty); err != nil {
		return s, err
	}

	rows, err := DB.Query(
		"SELECT c.name, SUM(qty) AS qty FROM wh_cell_register wr JOIN cell c ON c.id=wr.cell_id WHERE wr.product_id=$1 GROUP BY c.name",
		productId)

	if err != nil {
		return s, err
	}

	defer rows.Close()

	cells := []Cells{}

	for rows.Next() {
		var ce Cells
		if err := rows.Scan(&ce.CellName, &ce.Qty); err != nil {
			return s, err
		}
		cells = append(cells, ce)
	}
	s.Cells = cells

	return s, nil
}
