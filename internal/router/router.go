package router

import (
	"net/http"
	"transactionx/internal/handler"

	"github.com/gorilla/mux"
)

func InstanceRoutes(h handler.Handler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", h.HomePage()).Methods(http.MethodGet)
	router.HandleFunc("/index", h.FrontPage()).Methods(http.MethodGet)
	router.HandleFunc("/transactions", h.RegisterTransaction()).Methods(http.MethodPost)
	router.HandleFunc("/transactions", h.ListTransactions()).Methods(http.MethodGet)
	router.HandleFunc("/convert/{id}", h.ConvertTransaction()).Methods(http.MethodGet)

	return router
}
