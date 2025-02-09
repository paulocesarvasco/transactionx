package handler

import (
	"encoding/json"
	"net/http"
)

type Handler interface {
	HomePage() http.HandlerFunc
	RegisterTransaction() http.HandlerFunc
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to transactionX API!"})
	}
}

func (h *handler) RegisterTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
