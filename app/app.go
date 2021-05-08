// app.go

package app

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"fmt"
	"wms/app/controller"
)

type App struct {
	Router *mux.Router
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		// w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.Router.Use(CORS)
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/cells", controller.GetCells).Methods("GET")
	a.Router.HandleFunc("/cell", controller.CreateCell).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/cell/{id:[0-9]+}", controller.DeleteCell).Methods("DELETE", "OPTIONS")

	a.Router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product", controller.CreateProduct).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/productsearch/", controller.SearchProduct).Methods("GET")

	a.Router.HandleFunc("/receipts", controller.GetReceipts).Methods("GET")
	a.Router.HandleFunc("/receipt", controller.CreateReceipt).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/receipt/{id:[0-9]+}", controller.GetReceipt).Methods("GET")
	a.Router.HandleFunc("/receiptrow", controller.CreateReceiptRow).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/receiptrow/{id:[0-9]+}", controller.DeleteReceiptRow).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/receipt/{id:[0-9]+}", controller.ChangeStatusReceipt).Methods("PATCH", "OPTIONS")

	a.Router.HandleFunc("/allocations", controller.GetAllocations).Methods("GET")
	a.Router.HandleFunc("/allocation", controller.CreateAllocation).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/allocation/{id:[0-9]+}", controller.GetAllocation).Methods("GET")
	a.Router.HandleFunc("/allocation/{id:[0-9]+}", controller.UpdateAllocation).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc("/allocationrow", controller.CreateAllocationRow).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/allocationrow/{id:[0-9]+}", controller.UpdateAllocationRow).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc("/allocation/{id:[0-9]+}", controller.DeleteAllocation).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/allocationrows/{id:[0-9]+}", controller.DeleteAllocationRows).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/allocation/{id:[0-9]+}", controller.ChangeStatusAllocation).Methods("PATCH", "OPTIONS")

	a.Router.HandleFunc("/collects", controller.GetCollects).Methods("GET")
	a.Router.HandleFunc("/collect", controller.CreateCollect).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/collect/{id:[0-9]+}", controller.GetCollect).Methods("GET")
	a.Router.HandleFunc("/collectrow", controller.CreateCollectRow).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/collectrow/{id:[0-9]+}", controller.UpdateCollectRow).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc("/collectrow/{id:[0-9]+}", controller.DeleteCollectRow).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/collect/{id:[0-9]+}", controller.ChangeStatusCollect).Methods("PATCH", "OPTIONS")

	a.Router.HandleFunc("/waybills", controller.GetWaybills).Methods("GET")
	a.Router.HandleFunc("/waybill", controller.CreateWaybill).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/waybill/{id:[0-9]+}", controller.GetWaybill).Methods("GET")
	a.Router.HandleFunc("/waybill/{id:[0-9]+}", controller.UpdateWaybill).Methods("PUT", "OPTIONS")
	a.Router.HandleFunc("/waybillrow", controller.CreateWaybillRow).Methods("POST", "OPTIONS")
	a.Router.HandleFunc("/waybill/{id:[0-9]+}", controller.DeleteWaybill).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/waybillrows/{id:[0-9]+}", controller.DeleteWaybillRows).Methods("DELETE", "OPTIONS")
	a.Router.HandleFunc("/waybill/{id:[0-9]+}", controller.ChangeStatusWaybill).Methods("PATCH", "OPTIONS")

	a.Router.HandleFunc("/stock/{id:[0-9]+}", controller.GetStock).Methods("GET")
}
