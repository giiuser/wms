package controller

import (
	"database/sql"
	"fmt"
	"strconv"
	"wms/app/model"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetStock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := model.GetStocks(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Stock not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	} else {
		fmt.Println(err)
	}

	respondWithJSON(w, http.StatusOK, product)
}
