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

func GetAllocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid allocation ID")
		return
	}

	product, err := model.GetAllocation(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Allocation not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	} else {
		fmt.Println(err)
	}

	respondWithJSON(w, http.StatusOK, product)
}

func CreateAllocation(w http.ResponseWriter, r *http.Request) {
	var a model.Allocation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p, err := model.CreateAllocation(a.DocumentId, a.DocumentType)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func UpdateAllocation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var a model.Allocation
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	a.ID = id

	if err := model.UpdateAllocation(a.ID, a.DocumentId, a.DocumentType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, a)
}

func CreateAllocationRow(w http.ResponseWriter, r *http.Request) {
	var a model.AllocationRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateAllocationRow(a.AllocationId, a.ProductId, a.Qty, a.CellId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func UpdateAllocationRow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var ar model.AllocationRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ar); err != nil {
		fmt.Println(err)
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	ar.ID = id

	if err := model.UpdateAllocationRow(ar.ID, ar.CellId); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ar)
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
