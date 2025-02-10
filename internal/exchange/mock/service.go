package mock

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func NewServer(code int, payload any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		if payload != nil {
			json.NewEncoder(w).Encode(payload)
		}
	}))
}
