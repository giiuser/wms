package model

type Allocation struct {
	ID           int    `json:"id"`
	Status       int    `json:"status"`
	DocumentId   int    `json:"document_id"`
	DocumentType string `json:"document_type"`
}

type AllocationRow struct {
	AllocationId int `json:"allocation_id"`
	ProductId    int `json:"product_id"`
	Qty          int `json:"qty"`
}

func GetAllocation(id int) (Allocation, error) {
	var a Allocation
	row := DB.QueryRow("SELECT id, name FROM allocation WHERE id=$1", id)

	if err := row.Scan(&a.ID, &a.Status); err != nil {
		return a, err
	}

	return a, nil
}

func UpdateAllocation(id int, name string) error {
	_, err := DB.Exec("UPDATE allocation SET name=$1 WHERE id=$2",
		name, id)

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

func CreateAllocationRow(waybillId int, productId int, qty int) (AllocationRow, error) {
	var ar AllocationRow

	err := DB.QueryRow(
		"INSERT INTO allocation_table(allocation_id, product_id, qty) VALUES($1, $2, $3) RETURNING allocationl_id, product_id, qty",
		waybillId, productId, qty).Scan(&ar.AllocationId, &ar.ProductId, &ar.Qty)

	if err != nil {
		return ar, err
	}

	return ar, nil
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
