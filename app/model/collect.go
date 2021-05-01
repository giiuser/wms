// model.go

package model

type Collect struct {
	ID     int `json:"id"`
	Status int `json:"status"`
}

type CollectRow struct {
	CollectId int `json:"collect_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
}

func GetCollect(id int) (Collect, error) {
	var c Collect
	row := DB.QueryRow("SELECT id, name FROM collect WHERE id=$1", id)

	if err := row.Scan(&c.ID, &c.Status); err != nil {
		return c, err
	}

	return c, nil
}

func UpdateCollect(id int, name string) error {
	_, err := DB.Exec("UPDATE collect SET name=$1 WHERE id=$2",
		name, id)

	return err
}

func DeleteCollect(id int) error {
	_, err := DB.Exec("DELETE FROM collect WHERE id=$1", id)

	return err
}

func CreateCollect() (Collect, error) {
	var c Collect

	err := DB.QueryRow(
		"INSERT INTO receipt DEFAULT VALUES RETURNING id, status").Scan(&c.ID, &c.Status)

	if err != nil {
		return c, err
	}

	return c, nil
}

func CreateCollectRow(collectId int, productId int, qty int) (CollectRow, error) {
	var cr CollectRow

	err := DB.QueryRow(
		"INSERT INTO collect_table(collect_id, product_id, qty) VALUES($1, $2, $3) RETURNING collect_id, product_id, qty",
		collectId, productId, qty).Scan(&cr.CollectId, &cr.ProductId, &cr.Qty)

	if err != nil {
		return cr, err
	}

	return cr, nil
}

func ChangeStatusCollect(id int, status int) error {
	_, err := DB.Exec("UPDATE collect SET status=$1 WHERE id=$2",
		status, id)

	if status == 2 {
		MakeCellPosting(id, "collect", true)
	} else if status == 1 {
		MakeCellPosting(id, "collect", false)
	}

	return err
}
