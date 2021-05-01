// app.go

package app

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"wms/app/controller"
)

type App struct {
	Router *mux.Router
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// fmt.Println("ok")

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
	a.Router.HandleFunc("/products", controller.GetProducts).Methods("GET")
	a.Router.HandleFunc("/product", controller.CreateProduct).Methods("POST")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.GetProduct).Methods("GET")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.UpdateProduct).Methods("PUT")
	a.Router.HandleFunc("/product/{id:[0-9]+}", controller.DeleteProduct).Methods("DELETE")

	a.Router.HandleFunc("/receipt", controller.CreateReceipt).Methods("POST")
	a.Router.HandleFunc("/receiptrow", controller.CreateReceiptRow).Methods("POST")
	a.Router.HandleFunc("/receipt/{id:[0-9]+}", controller.ChangeStatusReceipt).Methods("PATCH")

	a.Router.HandleFunc("/waybill", controller.CreateWaybill).Methods("POST")
	a.Router.HandleFunc("/waybillrow", controller.CreateWaybillRow).Methods("POST")
	a.Router.HandleFunc("/waybill/{id:[0-9]+}", controller.ChangeStatusWaybill).Methods("PATCH")

	a.Router.HandleFunc("/allocation", controller.CreateAllocation).Methods("POST")
	a.Router.HandleFunc("/allocationrow", controller.CreateAllocationRow).Methods("POST")
	a.Router.HandleFunc("/allocation/{id:[0-9]+}", controller.ChangeStatusAllocation).Methods("PATCH")

	a.Router.HandleFunc("/collect", controller.CreateCollect).Methods("POST")
	a.Router.HandleFunc("/collectrow", controller.CreateCollectRow).Methods("POST")
	a.Router.HandleFunc("/collect/{id:[0-9]+}", controller.ChangeStatusCollect).Methods("PATCH")
}
