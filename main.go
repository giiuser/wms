// main.go

package main

import (
	"database/sql"
	"log"
	"wms/app"
	"wms/app/model"
)

func main() {
	var err error
	model.DB, err = sql.Open("postgres", "postgres://postgres:1234@localhost:5432/wms?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	a := app.App{}
	a.Initialize()

	a.Run(":8010")
}
