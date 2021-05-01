//wh_register.go

package model

import (
	"fmt"
)

type WhRegister struct {
	ID           int    `json:"id"`
	ProductId    int    `json:"product_id"`
	Qty          int    `json:"qty"`
	DocumentId   int    `json:"document_id"`
	DocumentType string `json:"document_type"`
}

type product struct {
	productId int
	qty       int
}

func MakePosting(documentId int, documentType string, direction bool) error {
	query := fmt.Sprintf("SELECT product_id, qty FROM %s_table WHERE %s_id=%d", documentType, documentType, documentId)
	rows, err := DB.Query(query)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.productId, &p.qty); err != nil {
			fmt.Println(err)
			return err
		}
		qty := p.qty
		if !direction {
			qty *= -1
		}
		createRow(p.productId, qty, documentId, documentType)
	}

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
