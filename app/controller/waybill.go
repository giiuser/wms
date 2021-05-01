package controller

import (
	"strconv"
	"wms/app/model"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateWaybill(w http.ResponseWriter, r *http.Request) {
	var wb model.Waybill
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wb); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p, err := model.CreateWaybill(wb.DocumentId, wb.DocumentType)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func CreateWaybillRow(w http.ResponseWriter, r *http.Request) {
	var wb model.WaybillRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wb); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateWaybillRow(wb.WaybillId, wb.ProductId, wb.Qty)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func ChangeStatusWaybill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var wb model.Waybill
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wb); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	wb.ID = id

	if err := model.ChangeStatusWaybill(wb.ID, wb.Status); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wb)
}
