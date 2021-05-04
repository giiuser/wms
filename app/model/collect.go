// model.go

package model

import "time"

type Collect struct {
	ID        int       `json:"id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type CollectRow struct {
	CollectId int `json:"collect_id"`
	ProductId int `json:"product_id"`
	Qty       int `json:"qty"`
	CellId    int `json:"cell_id"`
}

func GetCollect(id int) (Collect, error) {
	var c Collect
	row := DB.QueryRow("SELECT id, status, created_at FROM collect WHERE id=$1", id)

	if err := row.Scan(&c.ID, &c.Status, &c.CreatedAt); err != nil {
		return c, err
	}

	return c, nil
}

func GetCollects(start, count int) ([]Collect, error) {
	rows, err := DB.Query(
		"SELECT id, status, created_at FROM collect LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	collects := []Collect{}

	for rows.Next() {
		var c Collect
		if err := rows.Scan(&c.ID, &c.Status, &c.CreatedAt); err != nil {
			return nil, err
		}
		collects = append(collects, c)
	}

	return collects, nil
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
		"INSERT INTO collect DEFAULT VALUES RETURNING id, status").Scan(&c.ID, &c.Status)

	if err != nil {
		return c, err
	}

	return c, nil
}

func CreateCollectRow(collectId int, productId int, qty int, cell_id int) (CollectRow, error) {
	var cr CollectRow

	err := DB.QueryRow(
		"INSERT INTO collect_table(collect_id, product_id, qty, cell_id) VALUES($1, $2, $3, $4) RETURNING collect_id, product_id, qty, cell_id",
		collectId, productId, qty, cell_id).Scan(&cr.CollectId, &cr.ProductId, &cr.Qty, &cr.CellId)

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

func DeleteCollectRow(id int) error {
	_, err := DB.Exec("DELETE FROM collect_table WHERE id=$1", id)

	return err
}
