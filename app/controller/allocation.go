package controller

import (
	"strconv"
	"wms/app/model"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateAllocation(w http.ResponseWriter, r *http.Request) {
	var a model.Allocation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p, err := model.CreateWaybill(a.DocumentId, a.DocumentType)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func CreateAllocationRow(w http.ResponseWriter, r *http.Request) {
	var a model.AllocationRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateWaybillRow(a.AllocationId, a.ProductId, a.Qty)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func ChangeStatusAllocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var a model.Allocation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	a.ID = id

	if err := model.ChangeStatusAllocation(a.ID, a.Status); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, a)
}
