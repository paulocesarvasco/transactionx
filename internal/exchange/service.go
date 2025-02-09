package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"transactionx/internal/constants"
	"transactionx/internal/resources"
)

const BASE_URL string = "https://api.fiscaldata.treasury.gov/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=record_date,country,currency,exchange_rate&filter=record_date:gt:%s,country:eq:%s"

type service struct {
	c *http.Client
}

type Service interface {
	CountryData(ctx context.Context, country string) (resources.CountryMetadata, error)
}

func NewService() Service {
	return &service{
		c: &http.Client{},
	}
}

func (s *service) CountryData(ctx context.Context, country string) (resources.CountryMetadata, error) {
	startTime := time.Now().AddDate(0, -6, 0).Format("2006-01-02")
	finalURL := fmt.Sprintf(BASE_URL, startTime, country)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, finalURL, nil)
	if err != nil {
		log.Println("ERROR: ", err)
		return resources.CountryMetadata{}, constants.ErrorExchangeCreateRequest
	}
	res, err := s.c.Do(req)
	if err != nil {
		log.Println("ERROR: ", err)
		return resources.CountryMetadata{}, constants.ErrorExchangeRequestAPI
	}
	if res.StatusCode != http.StatusOK {
		return resources.CountryMetadata{}, constants.ErrorExchangeRequestUnsuccessful
	}
	defer res.Body.Close()
	var responsePayload resources.ExchangeAPIPayload
	err = json.NewDecoder(res.Body).Decode(&responsePayload)
	if err != nil {
		log.Println("ERROR: ", err)
		return resources.CountryMetadata{}, constants.ErrorExchangeDecodeResponse
	}
	if len(responsePayload.Data) == 0 {
		return resources.CountryMetadata{}, constants.ErrorExchangeRequestWithoutResults
	}
	return responsePayload.Data[len(responsePayload.Data)-1], nil
}
