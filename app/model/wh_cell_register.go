//model.go

package model

import (
	"fmt"
)

type WhCellRegister struct {
	ID           int    `json:"id"`
	ProductId    int    `json:"product_id"`
	Qty          int    `json:"qty"`
	DocumentId   int    `json:"document_id"`
	DocumentType string `json:"document_type"`
}

type cellProduct struct {
	cellId    int
	productId int
	qty       int
}

func MakeCellPosting(documentId int, documentType string, direction bool) error {
	query := fmt.Sprintf("SELECT product_id, qty, cell FROM %s_table WHERE %s_id=%d", documentType, documentType, documentId)
	rows, err := DB.Query(query)

	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var p cellProduct
		if err := rows.Scan(&p.cellId, &p.productId, &p.qty); err != nil {
			fmt.Println(err)
			return err
		}
		qty := p.qty
		if !direction {
			qty *= -1
		}
		createCellRow(p.cellId, p.productId, qty, documentId, documentType)
	}

	return nil
}

func createCellRow(cellId int, productId int, qty int, documentId int, documentType string) error {
	var wcr WhCellRegister

	err := DB.QueryRow(
		"INSERT INTO wh_cell_register(product_id, qty, document_id, document_type) VALUES($1, $2, $3, $4) RETURNING id",
		cellId, productId, qty, documentId, documentType).Scan(&wcr.ID)

	if err != nil {
		return err
	}

	return nil
}
