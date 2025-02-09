package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
