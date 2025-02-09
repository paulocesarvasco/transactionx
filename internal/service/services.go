package service

import (
	"time"
	"transactionx/internal/constants"
	"transactionx/internal/database"
	"transactionx/internal/resources"
)

type Services interface {
	RegisterTransaction(t resources.Transaction) (resources.Transaction, error)
}

type service struct {
	db database.Client
}

func NewService(db database.Client) Services {
	return &service{
		db: db,
	}
}

func (s *service) RegisterTransaction(t resources.Transaction) (resources.Transaction, error) {
	if len(t.Description) > constants.MAX_DESCRIPTION_LEN {
		return resources.Transaction{}, constants.ErrorInvliadDescriptionLenght
	}
	if _, err := time.Parse(time.RFC3339, t.Date); err != nil {
		return resources.Transaction{}, constants.ErrorInvliadTimeFormat
	}
	if t.PurchaseAmount <= 0 {
		return resources.Transaction{}, constants.ErrorTransactionPurchaseAmount
	}
	return s.db.RegisterTransaction(t)
}
