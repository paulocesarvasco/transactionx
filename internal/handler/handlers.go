package handler

import (
	"encoding/json"
	"net/http"
	"transactionx/internal/resources"
	"transactionx/internal/service"

	"github.com/gorilla/mux"
)

type Handler interface {
	HomePage() http.HandlerFunc
	FrontPage() http.HandlerFunc
	RegisterTransaction() http.HandlerFunc
	ListTransactions() http.HandlerFunc
	ConvertTransaction() http.HandlerFunc
}

func NewHandler(s service.Services) Handler {
	return &handler{
		s: s,
	}
}

type handler struct {
	s service.Services
}

func (h *handler) HomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to transactionX API!"})
	}
}

func (h *handler) FrontPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	}
}

func (h *handler) RegisterTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestPayload resources.Transaction
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&requestPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resources.Error{ResponseCode: http.StatusBadRequest, Message: err.Error()})
			return
		}
		t, err := h.s.RegisterTransaction(requestPayload)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resources.Error{ResponseCode: http.StatusBadRequest, Message: err.Error()})
			return

		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	}
}

func (h *handler) ListTransactions() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := h.s.ListTransactions()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resources.Error{ResponseCode: http.StatusBadRequest, Message: err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(list)
	}
}

func (h *handler) ConvertTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		transactionID := vars["id"]

		country := r.URL.Query().Get("country")

		ct, err := h.s.ConvertTransaction(r.Context(), transactionID, country)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(resources.Error{ResponseCode: http.StatusInternalServerError, Message: err.Error()})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ct)
	}
}
