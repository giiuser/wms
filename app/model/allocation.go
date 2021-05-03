package model

type Allocation struct {
	ID           int             `json:"id"`
	Status       int             `json:"status"`
	DocumentId   int             `json:"document_id"`
	DocumentType string          `json:"document_type"`
	Products     []AllocationRow `json:"products"`
}

type AllocationRow struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	AllocationId int    `json:"allocation_id"`
	ProductId    int    `json:"product_id"`
	Qty          int    `json:"qty"`
	CellId       int    `json:"cell_id"`
}

func GetAllocation(id int) (Allocation, error) {
	var a Allocation
	row := DB.QueryRow("SELECT id, status, document_id, document_type FROM allocation WHERE id=$1", id)

	if err := row.Scan(&a.ID, &a.Status, &a.DocumentId, &a.DocumentType); err != nil {
		return a, err
	}

	rows, err := DB.Query(
		"SELECT at.id, p.name, at.allocation_id, at.product_id, at.qty, at.cell_id FROM allocation_table at JOIN product p ON p.id = at.product_id WHERE allocation_id=$1",
		id)

	if err != nil {
		return a, err
	}

	defer rows.Close()

	products := []AllocationRow{}

	for rows.Next() {
		var ar AllocationRow
		if err := rows.Scan(&ar.ID, &ar.Name, &ar.AllocationId, &ar.ProductId, &ar.Qty, &ar.CellId); err != nil {
			return a, err
		}
		products = append(products, ar)
	}
	a.Products = products

	return a, nil
}

func UpdateAllocation(id int, documentId int, documentType string) error {
	_, err := DB.Exec("UPDATE allocation SET document_id=$1, document_type=$2 WHERE id=$3",
		documentId, documentType, id)

	return err
}

func DeleteAllocation(id int) error {
	_, err := DB.Exec("DELETE FROM allocation WHERE id=$1", id)

	return err
}

func CreateAllocation(documentId int, documentType string) (Allocation, error) {
	var a Allocation

	err := DB.QueryRow(
		"INSERT INTO allocation(document_id, document_type) VALUES($1, $2) RETURNING id, status, document_id, document_type",
		documentId, documentType).Scan(&a.ID, &a.Status, &a.DocumentId, &a.DocumentType)

	if err != nil {
		return a, err
	}

	return a, nil
}

func CreateAllocationRow(waybillId int, productId int, qty int, cellId int) (AllocationRow, error) {
	var ar AllocationRow

	err := DB.QueryRow(
		"INSERT INTO allocation_table(allocation_id, product_id, qty, cell_id) VALUES($1, $2, $3, $4) RETURNING allocation_id, product_id, qty, cell_id",
		waybillId, productId, qty, cellId).Scan(&ar.AllocationId, &ar.ProductId, &ar.Qty, &ar.CellId)

	if err != nil {
		return ar, err
	}

	return ar, nil
}

func UpdateAllocationRow(id int, cellId int) error {
	_, err := DB.Exec("UPDATE allocation_table SET cell_id=$1 WHERE id=$2",
		cellId, id)

	return err
}

func ChangeStatusAllocation(id int, status int) error {
	_, err := DB.Exec("UPDATE allocation SET status=$1 WHERE id=$2",
		status, id)

	if status == 2 {
		MakeCellPosting(id, "allocation", false)
	} else if status == 1 {
		MakeCellPosting(id, "allocation", true)
	}

	return err
}
