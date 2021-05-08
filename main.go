// main.go

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"wms/app"
	"wms/app/model"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"
)

func main() {
	var err error
	port, err := getPort()
	if err != nil {
		log.Fatal(err)
	}
	dbURL := os.Getenv("DATABASE_URL")
	// model.DB, err = sql.Open("postgres", "postgres://postgres:1234@localhost:5432/wms?sslmode=disable")
	model.DB, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	a := app.App{}
	a.Initialize()

	a.Run(port)
}

func getPort() (string, error) {
	// the PORT is supplied by Heroku
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
