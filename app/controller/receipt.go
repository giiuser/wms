package controller

import (
	"database/sql"
	"fmt"
	"strconv"
	"wms/app/model"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid receipt ID")
		return
	}

	product, err := model.GetReceipt(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Receipt not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	} else {
		fmt.Println(err)
	}

	respondWithJSON(w, http.StatusOK, product)
}

func GetReceipts(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := model.GetReceipts(start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func CreateReceipt(w http.ResponseWriter, r *http.Request) {
	var rec model.Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateReceipt()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func CreateReceiptRow(w http.ResponseWriter, r *http.Request) {
	var rec model.ReceiptRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateReceiptRow(rec.ReceiptId, rec.ProductId, rec.Qty)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func DeleteReceiptRow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
		return
	}

	if err := model.DeleteReceiptRow(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func ChangeStatusReceipt(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var rec model.Receipt
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&rec); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	rec.ID = id

	if err := model.ChangeStatusReceipt(rec.ID, rec.Status); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, rec)
}
