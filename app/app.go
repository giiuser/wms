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

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
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
}
