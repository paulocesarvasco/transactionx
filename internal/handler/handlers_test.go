package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"transactionx/internal/constants"
	"transactionx/internal/database"
	"transactionx/internal/exchange"
	"transactionx/internal/exchange/mock"
	"transactionx/internal/resources"
	"transactionx/internal/service"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestFrontPage(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/index.html", nil)
	rr := httptest.NewRecorder()
	exchangeService := exchange.NewService(&http.Client{}, constants.TREASURY_API_URL)
	h := NewHandler(service.NewService(database.NewSQLiteClient(), exchangeService))
	h.FrontPage().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusMovedPermanently, rr.Code, "test get front page")
}

func TestHomePage(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1:8080/", nil)
	rr := httptest.NewRecorder()
	exchangeService := exchange.NewService(&http.Client{}, constants.TREASURY_API_URL)
	h := NewHandler(service.NewService(database.NewSQLiteClient(), exchangeService))
	h.HomePage().ServeHTTP(rr, req)

	type response struct {
		Message string `json:"message"`
	}
	var responsePayload response
	json.NewDecoder(rr.Body).Decode(&responsePayload)
	assert.Equal(t, response{Message: "Welcome to transactionX API!"}, responsePayload, "test get home page")
}

func TestRegisterTransaction(t *testing.T) {
	tt := []struct {
		name         string
		transaction  any
		expectedCode int
	}{
		{"valid transaction", resources.Transaction{Description: "foo transaction", Date: "2025-02-10 15:00:00", PurchaseAmount: 100.00}, http.StatusCreated},
		{"long description", resources.Transaction{Description: "description with more than fifty characters very long description", Date: "2025-02-10 15:00:00", PurchaseAmount: 100.00}, http.StatusBadRequest},
		{"invalid time format", resources.Transaction{Description: "foo transaction", Date: "2025-02-02", PurchaseAmount: 100.00}, http.StatusBadRequest},
		{"invalid amount", resources.Transaction{Description: "foo transaction", Date: "2025-02-10 15:00:00", PurchaseAmount: -100.00}, http.StatusBadRequest},
		{"invalid payload", resources.CountryMetadata{Country: "Brazil"}, http.StatusBadRequest},
	}
	for _, tc := range tt {
		rawBody, _ := json.Marshal(tc.transaction)
		req, _ := http.NewRequest(http.MethodPost, "http://1270.0.0.1/transaction", bytes.NewReader(rawBody))
		rr := httptest.NewRecorder()
		exchangeService := exchange.NewService(&http.Client{}, constants.TREASURY_API_URL)
		h := NewHandler(service.NewService(database.NewSQLiteClient(), exchangeService))
		h.RegisterTransaction().ServeHTTP(rr, req)
		assert.Equal(t, tc.expectedCode, rr.Code, tc.name)
	}
}

func TestFetchTransactionsList(t *testing.T) {
	rawBody, _ := json.Marshal(resources.Transaction{Description: "foo transaction", Date: "2025-02-10 15:00:00", PurchaseAmount: 100.00})
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/transactions", bytes.NewReader(rawBody))
	rr := httptest.NewRecorder()

	exchangeService := exchange.NewService(&http.Client{}, constants.TREASURY_API_URL)
	h := NewHandler(service.NewService(database.NewSQLiteClient(), exchangeService))
	h.RegisterTransaction().ServeHTTP(rr, req)

	req, _ = http.NewRequest(http.MethodGet, "http://127.0.0.1/transactions", nil)
	rr = httptest.NewRecorder()
	h.ListTransactions().ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "fetch list")
}

func TestConvertTransaction(t *testing.T) {
	tt := []struct {
		name             string
		transactionID    string
		treasuryCode     int
		trasuryResponse  resources.ExchangeAPIPayload
		country          string
		expectedCode     int
		expectedResponse any
	}{
		{"success conversion", "", http.StatusOK,
			resources.ExchangeAPIPayload{Data: []resources.CountryMetadata{{ExchangeRate: "2", Country: "Brazil", Currency: "Real"}}},
			"Brazil", http.StatusOK, resources.ConvertedTransaction{ConvertedAmount: 200.00},
		},
		{"transaction not found", "1", http.StatusOK,
			resources.ExchangeAPIPayload{Data: []resources.CountryMetadata{{ExchangeRate: "2", Country: "Brazil", Currency: "Real"}}},
			"Brazil", http.StatusNotFound, resources.Error{ResponseCode: http.StatusNotFound, Message: constants.ErrorTransactionNotFound.Error()},
		},
	}

	for _, tc := range tt {
		rawBody, _ := json.Marshal(resources.Transaction{Description: "foo transaction", Date: "2025-02-10 15:00:00", PurchaseAmount: 100.00})
		req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/transactions", bytes.NewReader(rawBody))
		rr := httptest.NewRecorder()

		fakeAPI := mock.NewServer(tc.treasuryCode, tc.trasuryResponse)
		h := NewHandler(service.NewService(database.NewSQLiteClient(), exchange.NewService(fakeAPI.Client(), fakeAPI.URL+"?%s,%s")))
		h.RegisterTransaction().ServeHTTP(rr, req)
		assert.Equal(t, http.StatusCreated, rr.Code, tc.name)
		var registeredTransaction resources.Transaction
		json.NewDecoder(rr.Body).Decode(&registeredTransaction)

		var id string
		if tc.transactionID != "" {
			id = tc.transactionID
		} else {
			id = registeredTransaction.ID
		}

		url := "http://127.0.0.1:8080/convert/" + id + "?country=" + tc.country
		req, _ = http.NewRequest(http.MethodGet, url, nil)
		r := mux.NewRouter()
		r.HandleFunc("/convert/{id}", h.ConvertTransaction())
		rr = httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, tc.expectedCode, rr.Code, tc.name)
		if rr.Code == http.StatusOK {
			var response resources.ConvertedTransaction
			json.Unmarshal(rr.Body.Bytes(), &response)
			expectedResponse := tc.expectedResponse.(resources.ConvertedTransaction)
			assert.Equal(t, expectedResponse.ConvertedAmount, response.ConvertedAmount, tc.name)
		} else {
			var response resources.Error
			json.Unmarshal(rr.Body.Bytes(), &response)
			expectedResponse := tc.expectedResponse.(resources.Error)
			assert.Equal(t, expectedResponse, response, tc.name)
		}
	}
}
