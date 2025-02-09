package service

import (
	"time"
	"transactionx/internal/constants"
	"transactionx/internal/database"
	"transactionx/internal/resources"

	"github.com/google/uuid"
)

type Services interface {
	RegisterTransaction(t resources.Transaction) (resources.Transaction, error)
	ListTransactions() ([]resources.Transaction, error)
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
