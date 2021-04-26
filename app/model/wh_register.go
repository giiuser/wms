//wh_register.go

package model

type WhRegister struct {
	ID           int    `json:"id"`
	ProductId    int    `json:"product_id"`
	Qty          int    `json:"qty"`
	DocumentId   int    `json:"document_id"`
	DocumentType string `json:"document_type"`
}

func MakePosting(documentid int, documentType string) error {
	return nil
}

func createRow(productId int, qty int, documentId int, documentType string) error {
	var wr WhRegister

	err := DB.QueryRow(
		"INSERT INTO wh_register(product_id, qty, document_id, document_type) VALUES($1, $2, $3, $4) RETURNING id",
		productId, qty, documentId, documentType).Scan(&wr.ID)

	if err != nil {
		return err
	}

	return nil
}
