package service

import (
	"context"
	"math"
	"strconv"
	"time"
	"transactionx/internal/constants"
	"transactionx/internal/database"
	"transactionx/internal/exchange"
	"transactionx/internal/resources"

	"github.com/google/uuid"
)

type Services interface {
	RegisterTransaction(t resources.Transaction) (resources.Transaction, error)
	ListTransactions() ([]resources.Transaction, error)
	ConvertTransaction(ctx context.Context, transactionID string, country string) (resources.ConvertedTransaction, error)
}

type service struct {
	db       database.Client
	exchange exchange.Service
}

func NewService(db database.Client, conversor exchange.Service) Services {
	return &service{
		db:       db,
		exchange: conversor,
	}
}

func (s *service) RegisterTransaction(t resources.Transaction) (resources.Transaction, error) {
	if len(t.Description) > constants.MAX_DESCRIPTION_LEN {
		return resources.Transaction{}, constants.ErrorInvliadDescriptionLenght
	}
	if _, err := time.Parse("2006-01-02 15:04:05", t.Date); err != nil {
		return resources.Transaction{}, constants.ErrorInvliadTimeFormat
	}
	if t.PurchaseAmount <= 0 {
		return resources.Transaction{}, constants.ErrorTransactionPurchaseAmount
	}
	t.ID = uuid.New().String()
	return s.db.RegisterTransaction(t)
}

func (s *service) ListTransactions() ([]resources.Transaction, error) {
	return s.db.RetrieveTransactions()
}

func (s *service) ConvertTransaction(ctx context.Context, transactionID string, country string) (resources.ConvertedTransaction, error) {
	t, err := s.db.SearchTransaction(transactionID)
	if err != nil {
		return resources.ConvertedTransaction{}, err
	}
	cd, err := s.exchange.CountryData(ctx, country)
	if err != nil {
		return resources.ConvertedTransaction{}, err
	}
	var ct resources.ConvertedTransaction
	ct.Transaction = t
	ct.Currency = cd.Currency
	rate, _ := strconv.ParseFloat(cd.ExchangeRate, 64)
	rate = math.Round(rate*100) / 100
	ct.ExchangeRate = rate
	ct.ConvertedAmount = math.Round(rate*t.PurchaseAmount*100) / 100
	return ct, nil
}
