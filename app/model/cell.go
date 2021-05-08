// model.go

package model

type Cell struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func DeleteCell(id int) error {
	_, err := DB.Exec("DELETE FROM cell WHERE id=$1", id)

	return err
}

func CreateCell(name string) (Cell, error) {
	var c Cell
	err := DB.QueryRow(
		"INSERT INTO cell(name) VALUES($1) RETURNING id, name",
		name).Scan(&c.ID, &c.Name)

	if err != nil {
		return c, err
	}

	return c, nil
}

func GetCells(start, count int) ([]Cell, error) {
	rows, err := DB.Query(
		"SELECT id, name FROM cell LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cells := []Cell{}

	for rows.Next() {
		var c Cell
		if err := rows.Scan(&c.ID, &c.Name); err != nil {
			return nil, err
		}
		cells = append(cells, c)
	}

	return cells, nil
}
