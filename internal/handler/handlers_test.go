package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactionx/internal/resources"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/", nil)
	rr := httptest.NewRecorder()
	h := NewHandler()
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
		{"valid transaction", resources.Transaction{Description: "foo transaction", Date: "2025-02-02T12:00:00", PruchaseAmount: 100.00}, http.StatusCreated},
		{"invalid time format", resources.Transaction{Description: "foo transaction", Date: "2025-02-02", PruchaseAmount: 100.00}, http.StatusBadRequest},
		{"invalid amount", resources.Transaction{Description: "foo transaction", Date: "2025-02-02T12:00:00", PruchaseAmount: -100.00}, http.StatusBadRequest},
	}
	for _, tc := range tt {
		rawBody, _ := json.Marshal(tc.transaction)
		req, _ := http.NewRequest(http.MethodPost, "http://1270.0.0.1/transaction", bytes.NewReader(rawBody))
		rr := httptest.NewRecorder()
		h := NewHandler()
		h.HomePage().ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, tc.expectedCode, tc.name)
	}
}
