package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Db struct {
	DB *sql.DB
}

func (a *Db) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
