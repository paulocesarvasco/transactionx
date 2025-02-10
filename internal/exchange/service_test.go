package exchange

import (
	"context"
	"net/http"
	"testing"
	"transactionx/internal/constants"
	"transactionx/internal/exchange/mock"
	"transactionx/internal/resources"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	tt := []struct {
		name             string
		url              string
		serverCode       int
		serverResponse   any
		expectedResponse any
		expectedError    error
	}{
		{"failed create request", "", http.StatusOK, nil, resources.CountryMetadata{}, constants.ErrorExchangeCreateRequest},
		{"received not 200", "?filter=%s,country=%s", http.StatusInternalServerError, nil, resources.CountryMetadata{}, constants.ErrorExchangeRequestUnsuccessful},
		{"invalid response", "?filter=%s,country=%s", http.StatusOK, `{"invalid":"json"}`, resources.CountryMetadata{}, constants.ErrorExchangeDecodeResponse},
		{"response without results", "?filter=%s,country=%s", http.StatusOK, resources.ExchangeAPIPayload{}, resources.CountryMetadata{}, constants.ErrorExchangeRequestWithoutResults},
		{"all ok", "?filter=%s,country=%s", http.StatusOK, resources.ExchangeAPIPayload{Data: []resources.CountryMetadata{{Country: "Brazil"}}}, resources.CountryMetadata{Country: "Brazil"}, nil},
	}

	for _, tc := range tt {
		fakeServer := mock.NewServer(tc.serverCode, tc.serverResponse)
		s := NewService(fakeServer.Client(), fakeServer.URL+tc.url)
		response, err := s.CountryData(context.Background(), "Brazil")
		assert.Equal(t, tc.expectedError, err, tc.name)
		assert.Equal(t, tc.expectedResponse, response, tc.name)
	}
}
