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

func GetWaybill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid waybill ID")
		return
	}

	product, err := model.GetWaybill(id)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Waybill not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	} else {
		fmt.Println(err)
	}

	respondWithJSON(w, http.StatusOK, product)
}

func GetWaybills(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := model.GetWaybills(start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

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

func UpdateWaybill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid waybill ID")
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

	if err := model.UpdateWaybill(wb.ID, wb.DocumentId, wb.DocumentType); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, wb)
}

func DeleteWaybill(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Waybill ID")
		return
	}

	if err := model.DeleteWaybill(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteWaybillRows(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Allocation ID")
		return
	}

	if err := model.DeleteWaybillRows(id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
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
