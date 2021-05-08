package model

import (
	"database/sql"
	"time"
)

var DB *sql.DB

type BaseModel struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
