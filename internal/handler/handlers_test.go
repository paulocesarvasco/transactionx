package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"transactionx/internal/database"
	"transactionx/internal/resources"
	"transactionx/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
	rr := httptest.NewRecorder()
	h := NewHandler(service.NewService(database.NewSQLiteClient()))
	h.HomePage().ServeHTTP(rr, req)

	type response struct {
		Message string `json:"message"`
	}
	var responsePayload response
	json.NewDecoder(rr.Body).Decode(&responsePayload)
	assert.Equal(t, response{Message: "Welcome to transactionX API!"}, responsePayload)
}

func TestRegisterTransaction(t *testing.T) {
	tt := []struct {
		name         string
		transaction  resources.Transaction
		expectedCode int
	}{
		{"valid transaction", resources.Transaction{Description: "foo transaction", Date: time.Now().Format(time.RFC3339), PurchaseAmount: 100.00}, http.StatusCreated},
		{"invalid time format", resources.Transaction{Description: "foo transaction", Date: "2025-02-02", PurchaseAmount: 100.00}, http.StatusBadRequest},
		{"invalid amount", resources.Transaction{Description: "foo transaction", Date: "2025-02-02T12:00:00", PurchaseAmount: -100.00}, http.StatusBadRequest},
	}
	for _, tc := range tt {
		rawBody, _ := json.Marshal(tc.transaction)
		req, _ := http.NewRequest(http.MethodPost, "http://1270.0.0.1/transaction", bytes.NewReader(rawBody))
		rr := httptest.NewRecorder()
		h := NewHandler(service.NewService(database.NewSQLiteClient()))
		h.RegisterTransaction().ServeHTTP(rr, req)
		assert.Equal(t, tc.expectedCode, rr.Code, tc.name)
	}
}
