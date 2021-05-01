package controller

import (
	"strconv"
	"wms/app/model"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateCollect(w http.ResponseWriter, r *http.Request) {
	var c model.Collect
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&c); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()
	p, err := model.CreateCollect()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func CreateCollectRow(w http.ResponseWriter, r *http.Request) {
	var cr model.CollectRow
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&cr); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p, err := model.CreateCollectRow(cr.CollectId, cr.ProductId, cr.Qty)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func ChangeStatusCollect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		return

		var c model.Collect
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&c); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			return
		}
		defer r.Body.Close()
		c.ID = id

		if err := model.ChangeStatusWaybill(c.ID, c.Status); err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		respondWithJSON(w, http.StatusOK, c)
	}
}
