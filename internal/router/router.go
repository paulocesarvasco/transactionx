package router

import (
	"transactionx/internal/handler"

	"github.com/gorilla/mux"
)

func InstanceRoutes(h handler.Handler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", h.HomePage()).Methods("GET")

	return router
}
